package main

import (
	"bytes"
	"fmt"
	"my-go/day/20200110/longconn/sticky-bag/protocol"
	"net"
)

//粘包https://blog.csdn.net/weixin_33940102/article/details/93705145
var BUF_SIZE = 1024
var (
	buffer    = bytes.NewBuffer(make([]byte, 0, BUF_SIZE))
	readBytes = make([]byte, BUF_SIZE)
	isHead    = true
	bodyLen   = 0
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
		//tcpConn.SetReadDeadline(time.Now().Add(360 * time.Second))
		//tcpConn.SetWriteDeadline(time.Now().Add(360 * time.Second))
		if err != nil {
			continue
		}
		fmt.Println("wow  have a new client connected! ", tcpConn.RemoteAddr().String())
		go handleConn(tcpConn)

		go write(tcpConn)

	}

}
func write(conn *net.TCPConn) {
	words := "{\"id\":\" 2019 \",\"name\":\"golang\",\"message\":\"msg\"}\n"
	fmt.Println("server write data:", words)
	conn.Write(protocol.Packet([]byte(words)))
}

func handleConn(conn *net.TCPConn) {

	tempBuff := make([]byte, 0)
	readChannel := make(chan []byte, 16)

	go reader(readChannel)

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			//panic(err)
			fmt.Println("read client data error", err)
			break
		}
		tempBuff = protocol.UnPack(append(tempBuff, buff[:n]...), readChannel)
		fmt.Println("read data from client: ", string(tempBuff))
		write(conn)
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			fmt.Printf("server reader data:%s\n", string(data))
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
