package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	numOfCPU := runtime.NumCPU()
	fmt.Println("number of cpu: ", numOfCPU)

	runtime.GOMAXPROCS(numOfCPU)

	text := `
There are is an apple
A apple and a orange
read code and read code
write code and write code
`
	linesCh := readLines(text)
	wordsCh := splitWords(linesCh)
	result := countWords(wordsCh)
	fmt.Println("word count result: ", result)
}

func readLines(text string) <-chan string {
	lines := strings.Split(text, "\n")
	out := make(chan string, 1)
	go func() {
		for _, line := range lines {
			if line == "" || line == "\n" {
				continue
			}

			fmt.Println("[readLines] send line to lineCh: ", line)
			out <- line
		}
		close(out)
	}()

	return out
}

func splitWords(lineCh <-chan string) chan string {
	wordsCh := make(chan string, 1)

	go func() {
		for line := range lineCh {
			line = strings.TrimRight(line, "\n")
			words := strings.Split(line, " ")
			for _, word := range words {
				fmt.Printf("[splitWords] Sending word(%s) to words channel\n", word)
				wordsCh <- word
			}
		}
		close(wordsCh)
	}()

	return wordsCh
}

func countWords(wordsCh <-chan string) map[string]int {
	result := make(map[string]int, 10)
	for word := range wordsCh {
		fmt.Println("[countWords] Received word: ", word)
		result[word]++
	}

	return result
}
