package concurrentslice

import (
	"fmt"
	"sync"
)

// Iterator iter func
type Iterator func(v interface{}) bool

// Slice ...
type Slice interface {
	Append(v interface{}) error
	Range(Iterator) error
	Size() int
	Reset()
}

// SyncSlice thread safe slice
type SyncSlice struct {
	mu     sync.RWMutex
	values []interface{}
}

// NewSyncSlice size is the initial capacity of values slice
func NewSyncSlice(size int) Slice {
	if size <= 0 {
		panic(fmt.Errorf("invalid size"))
	}
	return &SyncSlice{
		values: make([]interface{}, 0, size),
	}
}

// Append append value to sync slice
func (s *SyncSlice) Append(v interface{}) error {
	s.mu.Lock()
	s.values = append(s.values, v)
	s.mu.Unlock()
	return nil
}

// Range if iter returns true then continue to loop over values
func (s *SyncSlice) Range(iter Iterator) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.values {
		if !iter(v) {
			return nil
		}
	}

	return nil
}

// Size the slice size
func (s *SyncSlice) Size() (size int) {
	s.mu.RLock()
	size = len(s.values)
	s.mu.RUnlock()
	return
}

// Reset reset values after all done
func (s *SyncSlice) Reset() {
	s.values = nil
}
