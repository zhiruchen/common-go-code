package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	timer := time.AfterFunc(10*time.Millisecond, func() {
		cancel()
	})

	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Printf("request get error: %v\n", err)
	}
	req = req.WithContext(ctx)

	fmt.Println("sending request...")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("sending err: ", err) // sending err:  Get "https://httpbin.org/get": context canceled
		return
	}

	defer resp.Body.Close()

	fmt.Println("reading body")
	for {
		timer.Reset(10 * time.Millisecond)
		_, err := io.CopyN(io.Discard, resp.Body, 256)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("read body err: ", err)
			return
		}
	}
}
