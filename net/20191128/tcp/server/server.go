package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("starting the server......")

	listener, err := net.Listen("tcp", "localhost:10000")

	if err != nil {
		fmt.Println("error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
			return
		}
		go doServerStuff(conn)

	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading", err.Error())
			return
		}
		fmt.Println("received data :%v", string(buf[:len]))
	}
}
