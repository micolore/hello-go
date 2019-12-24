package main

import "fmt"

func main() {
	fmt.Println("channel7 start...")
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	close(ch)
	<-ch //read
	fmt.Println("channel7 end...")
}
