package concurrentslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanSlice_Append(t *testing.T) {
	cases := []struct {
		desc         string
		values       []interface{}
		expectValues []interface{}
		err          error
	}{
		{
			desc:         "success append value",
			values:       []interface{}{1, 2, 3},
			expectValues: []interface{}{1, 2, 3},
			err:          nil,
		},
	}

	for _, cs := range cases {
		t.Run(cs.desc, func(t *testing.T) {
			s := NewChanSlice(len(cs.values))
			for _, v := range cs.values {
				if err := s.Append(v); err != nil {
					assert.Equal(t, cs.err, err)
				}
			}
			s.Close()

			var vs []interface{}
			iter := func(v interface{}) bool {
				vs = append(vs, v)
				return true
			}

			s.Range(iter)
			assert.Equal(t, cs.expectValues, vs)
		})
	}
}

func BenchmarkChanSliceAppend(b *testing.B) {
	for n := 1; n < b.N; n++ {
		s := NewChanSlice(n)
		for i := 1; i <= n; i++ {
			s.Append(i)
		}
		s.Close()
	}
}
