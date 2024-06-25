package mypackage

import (
    "bytes"
)

func maskLinks(message string) []byte {
    buffer := []byte(message)
    result := bytes.Buffer{}
    i := 0

    for i < len(buffer) {
        if len(buffer)-i >= 7 && isHttp(buffer[i:i+7]) {
            result.WriteString("[LINK REMOVED]")
            i += 7
            for i < len(buffer) && buffer[i] != ' ' {
                i++
            }
        } else {
            result.WriteByte(buffer[i])
            i++
        }
    }
    return result.Bytes()
}

func isHttp(sub []byte) bool {
    if len(sub) != 7 {
        return false
    }
    return (sub[0] == 'h' || sub[0] == 'H') &&
        (sub[1] == 't' || sub[1] == 'T') &&
        (sub[2] == 't' || sub[2] == 'T') &&
        (sub[3] == 'p' || sub[3] == 'P') &&
        (sub[4] == ':') &&
        (sub[5] == '/') &&
        (sub[6] == '/')
}
