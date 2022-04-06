package main

import (
	"context"
	`fmt`
	`net`

	"google.golang.org/grpc"

	pb "meta-grpc/grpc_test/proto"
)
const (
	address = "0.0.0.0:8080"
)
type Servers struct {}

func (s *Servers) SayHello(ctx context.Context, request *pb.HelloRequest)(*pb.HelloReply,error) {
	return &pb.HelloReply{
		Message: "hello" + request.Name,
	},nil
}

func main() {
	// 三部曲：
	//
	// 1、实例化gRPC的server
	r := grpc.NewServer()

	// 2、注册
	pb.RegisterGreeterServer(r,&Servers{})

	// 3、监听
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to Listening,err:%s",err)
	}
	err = r.Serve(lis)
	if err != nil {
		panic("Failed to start gRPC:"+err.Error())
	}
}
