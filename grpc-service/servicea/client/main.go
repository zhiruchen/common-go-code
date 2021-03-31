package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	svcA "github.com/zhiruchen/go-common/grpc-service/servicea/pb"
	"google.golang.org/grpc"
)

func main() {
	timeout := flag.Int("timeout", 100, "time out to service A")
	flag.Parse()

	if *timeout == 0 {
		*timeout = 50 // default 50ms
	}

	fmt.Println("timeout: ", *timeout, "ms")

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dail error: ", err)
		return
	}
	defer conn.Close()

	client := svcA.NewAServiceClient(conn)
	sayHelloWithTimeout(client, time.Duration(*timeout)*time.Millisecond)
}

func sayHelloWithTimeout(client svcA.AServiceClient, d time.Duration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	log.Printf("context: %+v\n", ctx)
	resp, err := client.SayHello(ctx, &svcA.SayHelloRequest{Msg: "A's Client"})
	if err != nil {
		log.Printf("[ASayHello] SayHello error: %v\n", err)
	} else {
		log.Printf("[ASayHello] SayHello: %v\n", resp.Resp)
	}

}
