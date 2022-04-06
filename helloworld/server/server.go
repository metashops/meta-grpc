package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

// Hello 方法是在HelloService 进行一层的封装方法
func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	// RPC 三部曲
	// 1、实例化一个server
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("Failed...", err.Error())
	}
	// 2、注册处理逻辑 Handler
	// HelloService 结构图注册到 RPC
	err = rpc.RegisterName("HelloServer", &HelloService{})
	if err != nil {
		fmt.Println("Failed...", err.Error())
	}

	// 3、启动服务
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed...", err.Error())
	}
	rpc.ServeConn(conn)
}
