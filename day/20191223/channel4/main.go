package main

import "fmt"

func Afunction(ch chan int) {
	fmt.Println("chennel4 finsh...")
	<-ch
}

func main() {
	fmt.Println("chennel4 start ...")

	ch := make(chan int, 10)
	for i := 0; i < 1000; i++ {
		ch <- 1
		go Afunction(ch)
	}
}
