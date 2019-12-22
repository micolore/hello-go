package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	fmt.Println("GrpcV2:HelloService:Hello:params: ", request, reply)
	*reply = "hello:" + request
	return nil
}

func main() {
	fmt.Println("grpc server language start.....")

	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}
