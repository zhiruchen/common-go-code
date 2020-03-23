package main

import "fmt"

// RangeNilMap ...
func RangeNilMap(m map[string]int) {
	for k, v := range m {
		fmt.Println("k: ", k, ", v: ", v)
	}
}

func main() {
	RangeNilMap(nil)
	RangeNilMap(map[string]int{"k1": 100, "k2": 200})
}
