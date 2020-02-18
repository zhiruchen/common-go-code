package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	path := "path_to_img/"
	name := "img_name.jpeg"
	if err := encode(path, name); err != nil {
		fmt.Println("encode error ", err)
	}
}

func encode(path, fileName string) error {
	f, err := os.Open(path + fileName)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	newFile, err := os.Create(path + "new_" + fileName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	return jpeg.Encode(newFile, img, &jpeg.Options{Quality: 30})
}
