package main

import (
	`fmt`
	`net`
	`sync`
	`time`

	`google.golang.org/grpc`

	pb `meta-grpc/stream_grpc_test/proto`
)

const (
	PROT = ":50052"
)
type server struct {}

// GetStream 服务端流模式：服务端返回一段连续的数据流
func (s *server)GetStream(req *pb.StreamReqData,res pb.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&pb.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// PutStream 客户端流模式
func (s *server)PutStream(cliStr pb.Greeter_PutStreamServer) error{
	for {
		if a,err := cliStr.Recv();err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

// AllStream 双向流模式
func (s *server)AllStream(allStr pb.Greeter_AllStreamServer) error {
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
			err := allStr.Send(&pb.StreamResData{
				Data: "我是服务器",
			})
			if err != nil {
				fmt.Println("Failed to send",err.Error())
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}


func main() {
	list,err := net.Listen("tcp",PROT)
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	pb.RegisterGreeterServer(g,&server{})

	err = g.Serve(list)
	if err != nil {
		fmt.Println(err)
	}

}
