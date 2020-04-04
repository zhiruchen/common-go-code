package concurrentslice

import (
	"fmt"
)

// ChanSlice thread safe, concurrent slice based on channel
type ChanSlice struct {
	ch chan interface{}
}

// NewChanSlice ...
func NewChanSlice(size int) *ChanSlice {
	if size <= 0 {
		panic(fmt.Errorf("invalid size"))
	}

	return &ChanSlice{
		ch: make(chan interface{}, size),
	}
}

// Append ...
func (s *ChanSlice) Append(v interface{}) (err error) {
	// the channel maybe closed
	defer func() {
		if e := recover(); e != nil {
			v, ok := e.(error)
			if ok {
				err = v
			}
		}
	}()

	// timer := time.NewTimer(time.Millisecond)
	// defer timer.Stop()

	select {
	case s.ch <- v:
	default:
	}
	return nil
}

// Range iter over values
func (s *ChanSlice) Range(iter Iterator) {
	for v := range s.ch {
		if !iter(v) {
			return
		}
	}
}

// Close close underlying channel
func (s *ChanSlice) Close() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	close(s.ch)
	return err
}

// Size size of values channel
func (s *ChanSlice) Size() int {
	return len(s.ch)
}
