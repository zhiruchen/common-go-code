package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	chatPb "github.com/zhiruchen/go-common/grpc/pb"
	"google.golang.org/grpc"
)

var (
	visitorMap = map[int32]*chatPb.Visitor{
		1: {Id: 1, Name: "visitor1"},
		2: {Id: 2, Name: "visitor2"},
	}

	messages = []*chatPb.ChatMessage{
		{
			Id:      1,
			FromId:  1,
			Content: "Hello, haha",
		},
		{
			Id:      1,
			FromId:  2,
			Content: "你好",
		},
		{
			Id:      1,
			FromId:  2,
			Content: "很好",
		},
		{
			Id:      1,
			FromId:  1,
			Content: "Good",
		},
		{
			Id:      1,
			FromId:  2,
			Content: "Great",
		},
	}
)

type chatServiceImpl struct{}

// GetVisitor(context.Context, *GetVisitorRequest) (*GetVisitorResponse, error)
//	GetMessage(ChatService_GetMessageServer) error

func (svc *chatServiceImpl) GetVisitor(ctx context.Context, request *chatPb.GetVisitorRequest) (resp *chatPb.GetVisitorResponse, err error) {
	visitor, ok := visitorMap[request.VisitorId]
	if !ok {
		return nil, fmt.Errorf("visitor not found")
	}
	return &chatPb.GetVisitorResponse{
		Visitor: visitor,
	}, nil
}

func (svc *chatServiceImpl) GetMessage(stream chatPb.ChatService_GetMessageServer) error {
	var msgCount int
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.Send(&chatPb.GetMessageResponse{})
		}
		if err != nil {
			return err
		}

		if msgCount >= len(messages) {
			return fmt.Errorf("message all send")
		}
		msg := messages[msgCount]
		msgCount++

		if err := stream.Send(&chatPb.GetMessageResponse{
			Messages: []*chatPb.ChatMessage{msg},
		}); err != nil {
			log.Printf("Error sending message to the client: %v\n", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5600")
	if err != nil {
		log.Printf("listen error: %v\n", err)
		return
	}

	s := grpc.NewServer()
	chatPb.RegisterChatServiceServer(s, &chatServiceImpl{})

	log.Println("-----------Start Chat Service On 5600---------")
	if err = s.Serve(lis); err != nil {
		log.Printf("serve error: %v\n", err)
		return
	}
}
