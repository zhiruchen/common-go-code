package main

import (
	"fmt"
	"strings"
	"time"
)

const chanBufferSize = 10

func main() {
	var (
		lineCh  = make(chan string, chanBufferSize)
		wordsCh = make(chan string, chanBufferSize)
	)

	text := `
There are is an apple
A apple and a orange
read code and read code
write code and write code
`
	go readLines(text, lineCh)
	go splitWords(lineCh, wordsCh)
	go countWords(wordsCh)
	time.Sleep(time.Second)
}

func readLines(text string, lineCh chan<- string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "" || line == "\n" {
			continue
		}

		fmt.Println("[readLines] send line to lineCh: ", line)
		lineCh <- line
	}
}

func splitWords(lineCh <-chan string, wordsCh chan<- string) {
	for {
		select {
		case line := <-lineCh:
			fmt.Println("[splitWords] received line: ", line)
			line = strings.TrimRight(line, "\n")
			words := strings.Split(line, " ")
			for _, word := range words {
				wordsCh <- word
			}
		default:
		}
	}
}

func countWords(wordsCh <-chan string) {
	for {
		select {
		case word := <-wordsCh:
			fmt.Println("[countWords] received word: ", word)
		default:
		}
	}
}
