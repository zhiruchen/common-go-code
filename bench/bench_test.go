package bench

import (
	"strconv"
	"sync"
	"testing"
)

var (
	number      = 1000000
	lockedStore = initStore(number, false)
	syncStore   = initSyncMap(number, false)
)

func initStore(n int, noValue bool) *LockedStore {
	s := &LockedStore{m: make(map[string]interface{}, n)}
	if noValue {
		return s
	}

	var i int
	for ; i < n; i++ {
		s.Set(strconv.Itoa(i), i)
	}
	return s
}

func initSyncMap(n int, noValue bool) sync.Map {
	var m sync.Map
	if noValue {
		return m
	}

	for i := 0; i < n; i++ {
		m.Store(strconv.Itoa(i), i)
	}
	return m
}

func BenchmarkSyncMapGet(b *testing.B) {
	n := b.N
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		k := strconv.Itoa(i)
		go func(key string) {
			syncStore.Load(key)
			wg.Done()
		}(k)
	}
	wg.Wait()
}

func BenchmarkSyncMapSet(b *testing.B) {
	n := b.N
	syncStore = initSyncMap(n, true)

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		k := strconv.Itoa(i)
		go func(key string) {
			syncStore.Store(key, ":"+key+":")
			wg.Done()
		}(k)
	}
	wg.Wait()
}

func BenchmarkLockedStore_Get(b *testing.B) {
	n := b.N
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		k := strconv.Itoa(i)
		go func(key string) {
			lockedStore.Get(key)
			wg.Done()
		}(k)
	}
	wg.Wait()
}

func BenchmarkLockedStore_Set(b *testing.B) {
	n := b.N
	lockedStore = initStore(n, true)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		k := strconv.Itoa(i)
		go func(key string) {
			lockedStore.Set(key, ":"+key+":")
			wg.Done()
		}(k)
	}
	wg.Wait()
}
