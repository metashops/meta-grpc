package main

import (
	`fmt`
	`net`
	`net/rpc`
	`net/rpc/jsonrpc`
)

func main() {
	conn,err := net.Dial("tcp","localhost:1234")
	if err != nil {
		panic("Failed to connection")
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello","body",&reply)
	if err != nil {
		fmt.Printf("Failed to Invoking,err:%s",err.Error())
	}
	fmt.Println(reply)
}
