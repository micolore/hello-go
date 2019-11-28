package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("first,what is you name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	for {
		fmt.Println("what to send to the server ? type q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + "says:" + trimmedInput))
	}

}
