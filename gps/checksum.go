// Package main provides connection to arduino serial.
package gps

import (
	"fmt"
	"strings"
	"encoding/hex"
	"encoding/binary"
)

func Checksum(gpsdata string) bool {
	
	data := strings.Split(gpsdata, "*")
	var checksum byte = 0;
        byteArr := []byte(data[0])
        for i := 1; i < len(data[0]); i++ {
                checksum = checksum ^ byteArr[i]
        }
        buf := make([]byte, 2)
        binary.LittleEndian.PutUint16(buf, uint16(checksum))
	if hex.EncodeToString(buf)[:2] == data[1] {
		return true
	}
	return false
}
