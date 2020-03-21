package main

import (
	"fmt"
	"sync"
)

func geValuesByKey(key string) []int {
	switch key {
	case "k1":
		return []int{1, 2, 3}
	case "k2":
		return []int{100, 200, 300}
	case "k6":
		return []int{8, 9, 10}
	}

	return []int{}
}

func concurrentGetValues(keys []string) map[string][]int {
	var valueStore sync.Map
	var wg sync.WaitGroup

	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			values := geValuesByKey(k)
			if len(values) > 0 {
				valueStore.Store(k, values)
			}
			wg.Done()
		}(key)
	}
	wg.Wait()

	result := make(map[string][]int, len(keys))
	valueStore.Range(func(key, value interface{}) bool {
		result[key.(string)] = value.([]int)
		return true
	})
	return result
}

func main() {
	keys := []string{"k1", "k2", "k6", "k3"}
	result := concurrentGetValues(keys)

	fmt.Printf("concurrentGetValues: %+v\n", result) // concurrentGetValues: map[k1:[1 2 3] k2:[100 200 300] k6:[8 9 10]]
}
