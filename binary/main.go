package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

const (
	payloadLen = 1
	sizeLen    = 4
	headerLen  = payloadLen + sizeLen
)

func main() {
	v := uint32(500)
	fmt.Println("v >> 24: ", v>>24)
	fmt.Println("v >> 16: ", v>>16)
	fmt.Println("v >> 8: ", v>>8)
	fmt.Println("v: ", v)
	fmt.Println("byte(v): ", byte(v))

	x := 'a'
	fmt.Println("byte(x): ", byte(x))
	fmt.Println("byte(100): ", byte(uint32(100)))
	n := uint32(302)
	fmt.Println("n: ", n, "byte(n): ", byte(n))

	hdr := make([]byte, headerLen)
	hdr[0] = byte(1)
	fmt.Println(hdr)
	binary.BigEndian.PutUint32(hdr[payloadLen:], uint32(500))
	fmt.Println("hdr: ", hdr)

	v = binary.BigEndian.Uint32(hdr[1:])
	fmt.Println(v)

	fmt.Println("encode data to buf")
	buf := encodeData()
	fmt.Println("encode buf: ", buf)
	fmt.Println("decode fields from buf")
	decodeBinaryData(buf)
}

func encodeData() []byte {
	buf := make([]byte, 10)
	ts := uint32(time.Now().Unix())

	fmt.Printf("encoding field1(buf[0:]) with: %x\n", 0xa20c)
	binary.BigEndian.PutUint16(buf[0:], 0xa20c)

	fmt.Printf("encoding field2(buf[2:]) with: %x\n", 0x04af)
	binary.BigEndian.PutUint16(buf[2:], 0x04af)

	fmt.Printf("encoding field2(buf[4:]) with: %d\n", ts)
	binary.BigEndian.PutUint32(buf[4:], ts)

	fmt.Printf("encoding field2(buf[8:]) with: %d\n", 888)
	binary.BigEndian.PutUint16(buf[8:], 888)

	return buf
}

func decodeBinaryData(buf []byte) {
	field1 := binary.BigEndian.Uint16(buf[0:])
	field2 := binary.BigEndian.Uint16(buf[2:])
	field3 := binary.BigEndian.Uint32(buf[4:])
	field4 := binary.BigEndian.Uint16(buf[8:])

	fmt.Printf("field1(buf[0:]): %x\n", field1)
	fmt.Printf("field2(buf[2:]): %x\n", field2)
	fmt.Printf("field3(buf[4:]): %d\n", field3)
	fmt.Printf("field4(buf[8:]): %d\n", field4)
}
