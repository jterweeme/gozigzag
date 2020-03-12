package gozigzag

import "fmt"

func Isin(c uint8, s string) bool {
    for _, chr := range s {
        if uint8(chr) == c {
            return true
        }
    }
    return false
}

func Isdigit(c uint8) bool {
    return Isin(c, "0123456789")
    return c >= '0' && c <= '9'
}

func Isxdigit(c uint8) bool {
    return Isin(c, "0123456789ABCDEFabcdef")
}

func Islower(c uint8) bool {
    return Isin(c, "abcdefghijklmnopqrstuvwxyz")
    return c >= 'a' && c <= 'z'
}

func Isupper(c uint8) bool {
    return Isin(c, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    return c >= 'A' && c <= 'Z'
}

func Isalpha(c uint8) bool {
    return Islower(c) || Isupper(c)
}

func Isalnum(c uint8) bool {
    return Isalpha(c) || Isdigit(c)
}

func Isspace(c uint8) bool {
    return c == ' '
}

func Ispunct(c uint8) bool {
    return Isin(c, "!\"#$%&'()*+`-.")
}

func Isprint(c uint8) bool {
    return Isalnum(c) || Ispunct(c)
}

func appendB(a *[]byte, b uint8) {
    *a = append(*a, byte(b))
}

func AppendBytes(a *[]byte, b uint8, n int) {
    for i := 0; i < n; i++ {
        appendB(a, b)
    }
}

func appendWLE(a *[]byte, w uint16) {
    *a = append(*a, byte(w >> 0 & 0xff))
    *a = append(*a, byte(w >> 8 & 0xff))
}

func appendDwLE(a *[]byte, dw uint32) {
    *a = append(*a, byte(dw >>  0 & 0xff))
    *a = append(*a, byte(dw >>  8 & 0xff))
    *a = append(*a, byte(dw >> 16 & 0xff))
    *a = append(*a, byte(dw >> 24 & 0xff))
}

func appendStr(a *[]byte, s string) {
    for _, chr := range s {
        *a = append(*a, byte(chr))
    }
}

func appendSlice(a *[]byte, b []byte) {
    for _, r := range b {
        *a = append(*a, r)
    }
}

func extendSlice(a *[]byte, n int) {
    length := len(*a)

    if (n < length) {
        fmt.Println("Error cannot extend slice!")
    }

    AppendBytes(a, 0, n - length)
}

func extractUint16LE(ba []byte, offset int) (uint16) {
    var ret uint16
    ret = uint16(ba[offset]) | uint16(ba[offset + 1]) << 8
    return ret
}

func extractUint32LE(ba []byte, offset int) (uint32) {
    var ret uint32
    ret = uint32(extractUint16LE(ba, offset))
    ret |= uint32(extractUint16LE(ba, offset + 2)) << 16
    return ret
}

func nibble(x uint8) (byte) {
    if x <= 9 {
        return byte(x) + '0'
    }
    if x <= 0xf {
        return byte(x) + 'a' - 10
    }
    return 'x'
}

func hex8(b uint8) string {
    var ret string
    ret += string(nibble(b >> 4 & 0xf))
    ret += string(nibble(b >> 0 & 0xf))
    return ret
}

func hex16(w uint16) string {
    var ret string
    ret += hex8(uint8(w >> 8 & 0xff))
    ret += hex8(uint8(w >> 0 & 0xff))
    return ret
}

func Hex32(dw uint32) string {
    var ret string
    ret += hex16(uint16(dw >> 16 & 0xffff))
    ret += hex16(uint16(dw >>  0 & 0xffff))
    return ret
}

func hexDump(ba []byte) {
    length := len(ba)

    for i := 0; i < length; i += 16 {
        var j int
        for j = 0; j < 16 && j + i < length; j++ {
            fmt.Print(hex8(ba[i + j]))
            fmt.Print(" ")
        }
        for j < 16 {
            fmt.Print("   ")
            j++
        }
        for j = 0; j < 16 && j + i < length; j++ {
            if Isprint(ba[i + j]) {
                fmt.Print(string(ba[i + j]))
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("\r\n")
    }
}

