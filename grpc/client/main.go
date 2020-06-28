package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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
	getVisitor(client)
	getVisitorMessages(client)
	getChatMessage(client)
	uploadVisitorAvatar(client)
}

func getVisitor(client chatPb.ChatServiceClient) {
	resp, err := client.GetVisitor(context.Background(), &chatPb.GetVisitorRequest{VisitorId: 1})
	if err != nil {
		log.Fatalf("[getVisitor] get visitor error: %v\n", err)
	}
	log.Printf("[getVisitor] get visitor: %v\n", resp.Visitor)
}

func getVisitorMessages(client chatPb.ChatServiceClient) {
	stream, err := client.GetVisitorMessage(context.Background(), &chatPb.GetVisitorMessageRequest{VisitorId: 1})
	if err != nil {
		fmt.Println("get message err: ", err)
	}

	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("get messages error: %v\n", err)
		}

		for _, msg := range result.Messages {
			log.Printf("[getVisitorMessages] get message: %+v\n", *msg)
		}
	}
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
				log.Printf("[getChatMessage] got message: %v", *in.Messages[0])
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

func uploadVisitorAvatar(client chatPb.ChatServiceClient) {
	stream, err := client.UploadVisitorAvatar(context.Background())
	if err != nil {
		log.Fatalf("[uploadVisitorAvatar] %v.UploadVisitorAvatar(_) = _, %v", client, err)
	}

	bs, err := ioutil.ReadFile("WechatIMG678.jpeg")
	if err != nil {
		log.Fatalf("[uploadVisitorAvatar] read file error: %v", err)
	}

	var batchSize = 2000
	fileLength := len(bs)
	total := fileLength / batchSize
	log.Printf("[uploadVisitorAvatar] fileLength: %d, totalCount: %d\n", fileLength, total)

	startIndex := 0
	for startIndex < total {
		start := startIndex * batchSize
		end := (startIndex + 1) * batchSize
		log.Printf("[uploadVisitorAvatar] start: %d, end: %d\n", start, end)
		if err := stream.Send(&chatPb.UploadVisitorAvatarRequest{Avatar: bs[start:end]}); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Send(bytes) = %v", stream, err)
		}
		startIndex++
	}

	if startIndex*batchSize < fileLength {
		if err := stream.Send(&chatPb.UploadVisitorAvatarRequest{Avatar: bs[startIndex*batchSize:]}); err != nil {
			log.Fatalf("%v.Send(bytes) = %v", stream, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	log.Printf("[uploadVisitorAvatar] resp: %v\n", reply)
}
