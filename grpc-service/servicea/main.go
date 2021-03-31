package main

import (
	"context"
	"fmt"
	"log"
	"net"

	svcA "github.com/zhiruchen/go-common/grpc-service/servicea/pb"
	svcB "github.com/zhiruchen/go-common/grpc-service/serviceb/pb"
	"google.golang.org/grpc"
)

type svcAImpl struct{}

func (svc *svcAImpl) SayHello(ctx context.Context, req *svcA.SayHelloRequest) (*svcA.SayHelloResponse, error) {
	log.Printf("[ServicceA] Incoming context: %+v\n", ctx)
	if err := callServiceBSayHello(ctx); err != nil {
		return nil, err
	}

	return &svcA.SayHelloResponse{Resp: "From A - Hello: " + req.Msg}, nil
}

func callServiceBSayHello(ctx context.Context) error {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dail error: ", err)
		return err
	}
	defer conn.Close()

	client := svcB.NewBServiceClient(conn)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	resp, err := client.SayHello(ctx, &svcB.SayHelloRequest{Msg: "FromA"})
	if err != nil {
		log.Printf("[BSayHello] SayHello error: %v\n", err)
		return err
	} else {
		log.Printf("[BSayHello] SayHello: %v\n", resp.Resp)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("listen error: %v\n", err)
		return
	}

	s := grpc.NewServer()
	svcA.RegisterAServiceServer(s, &svcAImpl{})

	log.Println("-----------Start ServiceA On 8080---------")
	if err = s.Serve(lis); err != nil {
		log.Printf("serve error: %v\n", err)
		return
	}
}
