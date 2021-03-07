package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestCallFnWithTimeout(t *testing.T) {
	t.Parallel()

	fnErr := fmt.Errorf("fn error")
	cases := []struct {
		desc    string
		fn      timeoutFn
		d       time.Duration
		wantErr error
	}{
		{
			desc: "call fn without timeout",
			fn: timeoutFn(func() error {
				fmt.Println("just return nil")
				return nil
			}),
			d:       100 * time.Millisecond,
			wantErr: nil,
		},
		{
			desc: "call fn has timeout",
			fn: timeoutFn(func() error {
				time.Sleep(110 * time.Millisecond)
				fmt.Println("consumed a longer time")
				return nil
			}),
			d:       100 * time.Millisecond,
			wantErr: errCallFnTimeout,
		},
		{
			desc: "return fn error",
			fn: timeoutFn(func() error {
				return fnErr
			}),
			d:       100 * time.Millisecond,
			wantErr: fnErr,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			err := CallFnWithTimeout(c.fn, c.d)
			if c.wantErr != err {
				t.Errorf("want: %v, got: %v\n", c.wantErr, err)
			}
		})
	}
}
