package main

import (
	"fmt"
	"my-go/day/20200110/longconn/sticky-bag/protocol"
	"net"
	"strconv"
	"time"
)

func main() {
	tcpAdd, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9090") //解析服务端TCP地址
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdd) //raddr是指远程地址，laddr是指本地地址，连接服务端
	//conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Println("net.DailTCP error:", err)
		return
	}
	go ReadMessage(conn)
	go WriteMessage(conn)
	time.Sleep(650 * time.Second)
}

func ReadMessage(conn *net.TCPConn) {

	tempBuff := make([]byte, 0)
	readChannel := make(chan []byte, 16)

	go reader(readChannel)

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			//panic(err)
			fmt.Println("from server read data error", err)
			break
		}
		tempBuff = protocol.UnPack(append(tempBuff, buff[:n]...), readChannel)
		fmt.Println("client read data:", string(tempBuff))
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			fmt.Printf("reader data:%s\n", string(data))
		}
	}
}

func WriteMessage(conn *net.TCPConn) {

	for i := 0; i < 2; i++ {
		id := strconv.Itoa(i)
		words := "{\"id\":\"" + id + "\",\"name\":\"golang\",\"message\":\"msg\"}\n"
		fmt.Println("client write data:", words)
		conn.Write(protocol.Packet([]byte(words)))
		time.Sleep(4 * time.Second)
	}
	fmt.Println("send data over!")
}
