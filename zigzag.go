package gozigzag

//import "fmt"
import "net"
import "io"

var g_n uint16

func response(conn *net.Conn) ([]byte) {
    var ba []byte
    AppendBytes(&ba, 0, 10)
    io.ReadFull(*conn, ba)
    length := extractUint32LE(ba, 6) + 5
    ba = nil
    extendSlice(&ba, int(length))
    io.ReadFull(*conn, ba)
    return ba
}

func macro2(ba []byte, n uint32) ([]byte) {
    var ret []byte
    appendWLE(&ret, 1)
    appendWLE(&ret, g_n)
    appendWLE(&ret, 1)
    appendDwLE(&ret, n)
    appendWLE(&ret, 0xfffd - (g_n + uint16(n)))
    appendSlice(&ret, ba)
    return ret
}

func keepalive(conn *net.Conn) {
    var ba []byte
    appendB(&ba, 0x16)
    ba = macro2(ba, 0)
    (*conn).Write(ba)
    response(conn)
}

func Login(conn *net.Conn) {
    g_n = 0xffff
    keepalive(conn)
    g_n = 0

    var ba []byte
    g_n = 0
    appendB(&ba, 1)
    appendStr(&ba, "SIGMATEK GmbH & Co KG")
    ba = macro2(ba, 0x15)
    (*conn).Write(ba)
    response(conn)
    
    ba = nil
    g_n = 2
    appendB(&ba, 0x71)
    ba = macro2(ba, 0)
    (*conn).Write(ba)
    response(conn)

    ba = nil
    g_n++
    appendB(&ba, 0x6d)
    ba = macro2(ba, 0)
    (*conn).Write(ba)
    response(conn)

    ba = nil
    g_n++
    appendWLE(&ba, 0x1b9)
    ba = macro2(ba, 1)
    (*conn).Write(ba)
    response(conn)
}

func GetAddr(conn *net.Conn, symbol string) uint32 {
    length := len(symbol) + 5
    var ba []byte
    appendB(&ba, 0x6c)
    appendWLE(&ba, uint16(length))
    appendB(&ba, 0)
    appendStr(&ba, symbol)
    appendB(&ba, 0)
    appendB(&ba, 8)
    ba = macro2(ba, uint32(length))
    (*conn).Write(ba)
    ba = response(conn)
    return extractUint32LE(ba, 11)
}

func ReadSvrUint(conn *net.Conn, addr uint32) uint32 {
    var ba []byte
    g_n++
    appendB(&ba, 0x6c)
    appendWLE(&ba, 0x0009)
    appendWLE(&ba, 0x0108)
    appendDwLE(&ba, addr)
    appendB(&ba, 8)
    ba = macro2(ba, uint32(9))
    (*conn).Write(ba)
    ba = response(conn)
    return extractUint32LE(ba, 11)
}





