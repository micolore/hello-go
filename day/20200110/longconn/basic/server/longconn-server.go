package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

// v1 接收客户端的消息，只需要读客户端传递端消息就可以了
func main() {

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	fmt.Println("server listen  bind ip:port 127.0.0.1:9090")
	tcpListenenr, _ := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("server listen start...")
	defer tcpListenenr.Close()

	for {
		tcpConn, err := tcpListenenr.AcceptTCP()
		//1秒超时读
		tcpConn.SetReadDeadline(time.Now().Add(360 * time.Second))
		tcpConn.SetWriteDeadline(time.Now().Add(360 * time.Second))
		if err != nil {
			continue
		}
		fmt.Println("wow  have a client connected! ", tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}
func tcpPipe(tcpConn *net.TCPConn) {
	fmt.Println("tcpPipe start...")
	ipStr := tcpConn.RemoteAddr().String()
	defer func() {
		fmt.Println("client disconnected :" + ipStr)
	}()
	for {
		message, err := ioutil.ReadAll(tcpConn)
		if err != nil {
			fmt.Println("read client  error!")
			panic(err)
		}
		if message != nil && len(message) > 0 {
			fmt.Println("client send message that is:", string(message))
		}
	}
}
