package main

import "fmt"

func main() {
	fmt.Println("channel2 start...")

	ch := make(chan int, 3)
	<-ch //刚创建是空的，没有东西可以取出

	fmt.Println("channel2 end....")
}
