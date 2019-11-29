package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {

	flag.Parse()
	if flag.NArg() != 2 {
		panic(("usage:host port"))
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		fmt.Println("hello .....")
		conn, err := listener.Accept()
		checkError(err, "Accept")
		go connectionHandler(conn)
	}

}
func connectionHandler(conn net.Conn) {

	connFrom := conn.RemoteAddr().String()
	println("connection form:", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONECT
		}
	}

DISCONECT:
	err := conn.Close()
	println("closed connection ", connFrom)
	checkError(err, "Close:")
}

func sayHello(conn net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	write, err := conn.Write(obuf)
	checkError(err, "write : write "+string(write)+"bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func initServer(hostAndPort string) net.Listener {

	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "resolving address:post failed:"+hostAndPort+"")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "listen tcp:")
	println("listening to:", listener.Addr().String())
	return listener

}

func checkError(err error, info string) {

	if err != nil {
		panic("error:" + info + " " + err.Error())
	}
}
