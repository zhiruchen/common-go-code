package main

import "fmt"

func main() {
	var hashTable = make(map[string]int, 10)
	fmt.Println("hashTable", hashTable, ", length: ", len(hashTable))
}
