package main

import "fmt"

func main() {
	fmt.Println("channel6 start...")
	ch := make(chan int, 5)
	ch <- 1
	ch <- 1
	close(ch)
	ch <- 1
	fmt.Println("channel6 end...")
}
