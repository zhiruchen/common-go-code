package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	svcB "github.com/zhiruchen/go-common/grpc-service/serviceb/pb"
	"google.golang.org/grpc"
)

type svcBImpl struct{}

func (svc *svcBImpl) SayHello(ctx context.Context, req *svcB.SayHelloRequest) (*svcB.SayHelloResponse, error) {
	fmt.Printf("context; %+v\n", ctx)

	fmt.Println("-----Service B consume 40ms----")
	time.Sleep(time.Duration(40) * time.Millisecond)
	fmt.Println("-----and then call Service C---")
	if err := callServiceC(ctx); err != nil {
		return nil, err
	}

	return &svcB.SayHelloResponse{Resp: "FromB - Hello: " + req.Msg}, nil
}

func callServiceC(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequest("GET", "http://localhost:8082", nil)
	if err != nil {
		log.Println("new req error: ", err)
		return err
	}
	req = req.WithContext(ctx)

	log.Println("B Sending request to Service C...")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Service C returns error: ", err)
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("get resp from Service C: ", string(bytes))
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("listen error: %v\n", err)
		return
	}

	s := grpc.NewServer()
	svcB.RegisterBServiceServer(s, &svcBImpl{})

	log.Println("-----------Start ServiceB On 8081---------")
	if err = s.Serve(lis); err != nil {
		log.Printf("serve error: %v\n", err)
		return
	}
}
