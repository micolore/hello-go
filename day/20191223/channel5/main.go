package main

import "fmt"

func Afunction(ch chan int) {

	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	<-ch
}

func main() {
	fmt.Println("channel5 start...")
	ch := make(chan int, 10)
	for i := 0; i < 100; i++ {
		ch <- 1
		go Afunction(ch)
	}

}
