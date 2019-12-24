package main

import (
	"fmt"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("reply")
}
