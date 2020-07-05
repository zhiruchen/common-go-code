package bench

import "sync"

type LockedStore struct {
	sync.RWMutex
	m map[string]interface{}
}

func (s *LockedStore) Get(key string) interface{} {
	s.RLock()
	defer s.RUnlock()
	return s.m[key]
}

func (s *LockedStore) Set(key string, value interface{}) {
	s.Lock()
	s.m[key] = value
	s.Unlock()
}
