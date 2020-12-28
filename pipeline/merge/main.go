package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) (out chan int) {
	out = make(chan int, len(nums))
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func merge(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	output := func(in <-chan int) {
		for n := range in {
			out <- n
		}
		wg.Done()
	}

	for _, ch := range chans {
		wg.Add(1)
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(3, 6, 8, 9, 5, 10, 12)
	out1 := sq(in)
	out2 := sq(in)

	for n := range merge(out1, out2) {
		fmt.Println("merge n: ", n)
	}
}
