package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:10000")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "localhost:10000")
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80")
	checkConnection(conn, err)
}
func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Println("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %d\n", conn)
}
