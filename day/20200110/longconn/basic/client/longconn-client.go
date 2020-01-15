package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAdd, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9090") //解析服务端TCP地址
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdd) //raddr是指远程地址，laddr是指本地地址，连接服务端
	conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	if err != nil {
		fmt.Println("net.DailTCP error:", err)
		return
	}
	go ReadMessage(conn)
	go WriteMessage(conn)
	time.Sleep(65 * time.Second)
}

func ReadMessage(conn *net.TCPConn) {

	fmt.Println("readMessage start...")

}
func WriteMessage(conn *net.TCPConn) {

	fmt.Println("writeMessage start...")
	b := []byte("hello long conn\n")
	r, err := conn.Write(b)
	if err != nil {
		fmt.Println("write error!")
	}
	fmt.Println("writeMessage first...", r)
	time.Sleep(20 * time.Second)

	b = []byte("hello long conn second!\n")
	r, err = conn.Write(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("writeMessage second...!", r)

	time.Sleep(10 * time.Second)
}
