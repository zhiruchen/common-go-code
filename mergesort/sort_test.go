package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	v1, v2 := []int{1, 3, 8}, []int{2, 5, 6}
	result := merge(v1, v2)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 8}, result)
}

func TestMergeSort(t *testing.T) {
	vs := []int{3, 2, 5, 0, 8, 1}
	result := MergeSort(vs)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 8}, result)
}
