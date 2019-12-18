package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test()
}

var vim struct {
	name string
}

func test() {

	fmt.Println("test")
}
