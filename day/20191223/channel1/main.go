package main

import "fmt"

func main() {
	fmt.Println("channel 1 start")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1 //channel 已经满了

	fmt.Println("channel 1 end ")
}
