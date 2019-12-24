package main

import (
	"fmt"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {

	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	fmt.Println("grpcv3.start......")
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Hello("hello grpc-v3", &reply)
	if err != nil {
		panic(err)
	}
}
