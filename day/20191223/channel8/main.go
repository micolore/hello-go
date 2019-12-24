package main

import "fmt"

func main() {
	fmt.Println("channel8 start...")
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("channel8 data:", data)
	}
	fmt.Println("channel8 end...")
}
