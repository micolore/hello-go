package main

import (
	"fmt"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

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

	RegisterHelloService(new(HelloService))
	fmt.Println("grpc server v2 start.....")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)

}
