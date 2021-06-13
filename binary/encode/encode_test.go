package encode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	data := []byte("Hello Binary")
	buf := EncodeDataToBytes(data)
	fmt.Println("buf content: ", buf)

	mode, payloadLen, msg := DdecodedataFromBytes(buf)
	fmt.Println("mode: ", mode)
	fmt.Println("payloadLen: ", payloadLen)
	fmt.Println("msg: ", msg)

	assert.Equal(t, uint8(1), mode)
	assert.Equal(t, uint32(12), payloadLen)
	assert.Equal(t, []byte("Hello Binary"), msg)
}
