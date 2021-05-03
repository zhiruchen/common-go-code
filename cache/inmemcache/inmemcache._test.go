package inmemcache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	defer kvstore.Flush()

	key1, value1 := "test-key1", "test-value1"
	Set(key1, value1, 100*time.Millisecond)
	v1, ok := Get(key1)
	assert.True(t, ok)
	assert.Equal(t, value1, v1)

	t.Log("waiting key1 to expire")
	time.Sleep(100 * time.Millisecond)
	v1, ok = Get(key1)
	fmt.Println("v1: ", v1, ", ok: ", ok)
	assert.False(t, ok)
	assert.Nil(t, v1)
}
