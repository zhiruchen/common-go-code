package main

import "fmt"

func main() {
	c := make(chan struct{})
	go func() {
		fmt.Println("Hello from goroutine")
		c <- struct{}{}
	}()

	<-c
	fmt.Println("vim-go")
}
