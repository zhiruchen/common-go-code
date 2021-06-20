package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		desc         string
		input        []int
		expectResult []int
		filterFn     func(n int) bool
	}{
		{
			desc:         "filter even number",
			input:        []int{1, 2, 3, 6, 8, 9, 110},
			expectResult: []int{2, 6, 8, 110},
			filterFn: func(n int) bool {
				return n%2 == 0
			},
		},
		{
			desc:         "filter odd number",
			input:        []int{1, 2, 3, 6, 8, 9, 110},
			expectResult: []int{1, 3, 9},
			filterFn: func(n int) bool {
				return n%2 != 0
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expectResult, Filter(tc.input, tc.filterFn))
		})
	}
}

func TestFilterInPlace(t *testing.T) {
	cases := []struct {
		desc         string
		input        []int
		expectResult []int
		filterFn     func(n int) bool
	}{
		{
			desc:         "filter even number",
			input:        []int{1, 2, 3, 6, 8, 9, 110},
			expectResult: []int{2, 6, 8, 110},
			filterFn: func(n int) bool {
				return n%2 == 0
			},
		},
		{
			desc:         "filter odd number",
			input:        []int{1, 2, 3, 6, 8, 9, 110},
			expectResult: []int{1, 3, 9},
			filterFn: func(n int) bool {
				return n%2 != 0
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expectResult, FilterInPlace(tc.input, tc.filterFn))
		})
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/zhiruchen/go-common/slice
cpu: VirtualApple @ 2.50GHz
BenchmarkFilter-8   	26934061	        43.56 ns/op	      64 B/op	       1 allocs/op
*/
func BenchmarkFilter(b *testing.B) {
	s := []int{1, 2, 3, 6, 8, 9, 110}
	for n := 0; n < b.N; n++ {
		_ = Filter(s, func(n int) bool {
			return n%2 == 0
		})
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/zhiruchen/go-common/slice
cpu: VirtualApple @ 2.50GHz
BenchmarkFilterInPlace-8   	60510556	        19.45 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkFilterInPlace(b *testing.B) {
	s := []int{1, 2, 3, 6, 8, 9, 110}
	for n := 0; n < b.N; n++ {
		_ = FilterInPlace(s, func(n int) bool {
			return n%2 == 0
		})
	}
}
