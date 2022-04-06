package main

import (
	`context`
	`fmt`

	"google.golang.org/grpc"

	pb"meta-grpc/grpc_test/proto"
)
const (
	address = "127.0.0.1:8080"
)
func main() {
	// 1、连接
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		panic("Failed to connection"+err.Error())
	}
	defer conn.Close()

	// 2、实例化一个Client
	c := pb.NewGreeterClient(conn)

	// 3、调用
	r,err := c.SayHello(context.Background(),&pb.HelloRequest{
		Name: " bobby",
	})
	if err != nil {
		panic("Failed to Invoking"+err.Error())
	}
	fmt.Println(r.Message)

}
