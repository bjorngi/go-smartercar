package gps

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

// Checksum checks if given GPS data is correct by checking the data between $ and * against
// the data after *. This is done with xor.
// It returns true or false whether or not the data is correct.
func Checksum(gpsdata string) bool {

	data := strings.Split(gpsdata, "*")
	var checksum byte
	byteArr := []byte(data[0])
	for i := 1; i < len(data[0]); i++ {
		checksum = checksum ^ byteArr[i]
	}
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(checksum)) //Not possible with uint8
	if hex.EncodeToString(buf)[:2] == data[1] {          //Therefore we use only the first 2 numbers
		return true
	}
	return false
}
