package main

import (
	"fmt"
	"net"
	"os"
)

/*
 * ip 191.168.1.191
 *
 */
func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("invalid address", addr.String())
	} else {
		fmt.Println("the address is ", addr.String())
	}
	os.Exit(0)
}
