package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) (out chan int) {
	out = make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case out <- num:
			case <-done: // a receive operation on a closed channel can always proceed immediately
				fmt.Println("[gen] Received done")
				return
			}

			out <- num
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-done:
				fmt.Println("[sq] Received done")
				return
			}
		}
	}()

	return out
}

func merge(done <-chan struct{}, chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	output := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			select {
			case out <- n:
			case <-done:
				fmt.Println("[merge] Received done")
				return
			}
		}
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

// main decide to exit without receiving all the values from out.
// it must tell the goroutines in the upstreams to abandon the values
// they're trying to send
func main() {
	done := make(chan struct{})
	defer close(done) // main unlock all senders by closing the done channel

	in := gen(done, 9, 10)
	out1, out2 := sq(done, in), sq(done, in)

	out := merge(done, out1, out2)
	fmt.Println(<-out) // only received one value from merged channel
}
