package main

import (
	`context`
	`fmt`
	`sync`
	`time`

	`google.golang.org/grpc`

	pb`meta-grpc/stream_grpc_test/proto`
)

func main() {
	// 一、服务端流模式
	//
	// 1、拨号，你可以理解，请求连接对方
	conn,err := grpc.Dial("localhost:50052",grpc.WithInsecure())
	if err != nil {
		panic("Failed to connection"+err.Error())
	}
	defer conn.Close()

	// 2、建立连接
	c := pb.NewGreeterClient(conn)
	res,_ := c.GetStream(context.Background(),&pb.StreamReqData{
		Data: "hallo",
	})
	for {
		recv, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Data)
	}

	// 二、客户端流模式
	//
	putS,_ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&pb.StreamReqData{
			Data: fmt.Sprintf("我是客户端发送流%d",i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 三、双向流模式
	//
	allStr,_ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data,_ := allStr.Recv()
			fmt.Println("收到客户端消息："+data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := allStr.Send(&pb.StreamReqData{
				Data: "我是客户端",
			})
			if err != nil {
				fmt.Println("Failed to send",err.Error())
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}
