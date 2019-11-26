package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("tcp server start.....")
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {

	defer conn.Close()

	dayTime := time.Now().String()
	fmt.Println("client request time :", dayTime)
	conn.Write([]byte("hello client"))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "tcp server fatal error: %s", err.Error())
		os.Exit(1)
	}
}
