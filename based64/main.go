package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func base64ToImgBytes(base64Str string) ([]byte, error) {
	unbased, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return unbased, nil
}

func main() {
	base64Str := ``
	imgBytes, err := base64ToImgBytes(base64Str)
	if err != nil {
		fmt.Print("based64 decode error: ", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("test_file", imgBytes, 0644)
	fmt.Println("write file error: ", err)
}
