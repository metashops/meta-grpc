package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8800")
	if err != nil {
		panic("Failed to connection")
	}
	var reply string
	err = client.Call("HelloService.Hello", "body", &reply)
	if err != nil {
		panic("Failed to Invoking")
	}
	fmt.Println(reply)
}
