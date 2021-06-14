package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	dirPath := flag.String("dir", "", "the dir path")
	flag.Parse()

	if *dirPath == "" {
		*dirPath = "."
	}

	fileSizeCh := make(chan int64)
	var wg sync.WaitGroup

	wg.Add(1)
	go calculateDirSize(*dirPath, &wg, fileSizeCh)

	go func() {
		wg.Wait()
		close(fileSizeCh)
	}()

	var fileNum, totalFileSize int64
	for size := range fileSizeCh {
		fileNum++
		totalFileSize += size
	}

	fmt.Printf("fileNum: %d, totalFileSize: %d\n", fileNum, totalFileSize)
}

func calculateDirSize(dir string, wg *sync.WaitGroup, fileSizeCh chan<- int64) {
	defer wg.Done()

	files := dirFiles(dir)
	for _, file := range files {
		fmt.Printf("isDir: %v, calculating name: %s\n", file.IsDir(), file.Name())

		if file.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, file.Name())
			go calculateDirSize(subDir, wg, fileSizeCh)
		} else {
			fileSizeCh <- file.Size()
		}
	}
}

// limit the goroutine count that enters dirFiles function
var sema = make(chan struct{}, 20)

func dirFiles(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get files in: %s, has error: %v\n", dir, err)
		return []os.FileInfo{}
	}

	return files
}
