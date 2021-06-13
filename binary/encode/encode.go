package encode

import (
	"encoding/binary"
)

func EncodeDataToBytes(msg []byte) []byte {
	payloadLength := 1
	sizeLen := 4

	length := len(msg)
	buf := make([]byte, payloadLength+sizeLen+length)

	buf[0] = byte(1)
	binary.BigEndian.PutUint32(buf[1:], uint32(length))
	copy(buf[5:], msg)

	return buf
}

func DdecodedataFromBytes(encodeData []byte) (mode uint8, payloadLen uint32, msg []byte) {
	mode = encodeData[0]
	payloadLen = binary.BigEndian.Uint32(encodeData[1:])
	msg = encodeData[5:]
	return
}
