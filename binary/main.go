package main

import (
	"encoding/binary"
	"fmt"
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
	fmt.Println(hdr)

	v = binary.BigEndian.Uint32(hdr[1:])
	fmt.Println(v)
}
