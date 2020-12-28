package main

import "fmt"

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

func main() {
	in := gen(2, 3)
	out := sq(in)
	fmt.Println("num1: ", <-out)
	fmt.Println("num2: ", <-out)
}
