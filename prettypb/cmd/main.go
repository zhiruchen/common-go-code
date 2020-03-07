package main

import (
	"flag"
	"fmt"
)

func main() {
	pbMsgFile := flag.String("f", "", "pb message file path")
	flag.Parse()

	if len(*pbMsgFile) == 0 {
		fmt.Println("need msg file")
		return
	}

}
