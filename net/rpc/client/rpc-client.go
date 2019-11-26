package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "192.168.1.191:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{80, 80}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d \n", args.A, args.B, reply)

	//异步
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done
	fmt.Printf("Arith.Divide %d", replyCall.Done)
}
