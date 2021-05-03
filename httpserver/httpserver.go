package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("run server error: %v\n", err)
	}
}
