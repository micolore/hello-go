package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

/*
 * tcp client fatal error :read tcp 192.168.1.191:52579->192.168.1.191:7777: read: connection reset by peer exit status 1
 *
 */
func main() {
	fmt.Println("tcp client start.....")
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	fmt.Println("tcp server return result ", string(result))
	checkError(err)
	os.Exit(0)
}

func checkError(err error) {

	if err != nil {
		fmt.Println(os.Stderr, "tcp client fatal error :%s ", err.Error())
		//os.Exit(1)
	}
}
