package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	fmt.Println("HelloService:Hello:params: ", request, reply)
	*reply = "hello:" + request
	return nil
}

func main() {

	rpc.RegisterName("HelloService", new(HelloService))

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
