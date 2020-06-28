package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	chatPb "github.com/zhiruchen/go-common/grpc/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:5600", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dail error: ", err)
		return
	}
	defer conn.Close()

	client := chatPb.NewChatServiceClient(conn)
	getChatMessage(client)
}

func getChatMessage(client chatPb.ChatServiceClient) {
	stream, err := client.GetMessage(context.Background())
	if err != nil {
		fmt.Println("get message err: ", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			if len(in.Messages) >= 1 {
				log.Printf("Got message: %v", *in.Messages[0])
			}
		}
	}()

	requests := []*chatPb.GetMessageRequest{
		{},
		{},
		{},
		{},
		{},
	}
	for _, req := range requests {
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
