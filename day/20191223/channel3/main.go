package main

import "fmt"

func Afunction(ch chan int) {
	fmt.Println("Afunction... finsh")
	<-ch
}

func main() {

	fmt.Println("channel3 start...")
	ch := make(chan int) //无缓冲
	go Afunction(ch)
	ch <- 1
	fmt.Println("channel3 end...")

}
