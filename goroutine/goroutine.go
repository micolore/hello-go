// goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello goroutine&channel!")

	//go f(x,y,z)

	//go say("world")
	//say("hello")

	//channel
	//ch <- v  (v发送到ch)
	//v := <-ch(ch接收数据,并赋值给v)
	//ch := make(chan int)

	s := []int{1, 8, 2, 3, 4, 5, 7}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func say(s string) {

	for x := 0; x < 10; x++ {
		go cal(x, x+1)
	}
	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func sum(s []int, c chan int) {

	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum //把sum发送给c

}

func cal(a, b int) {
	fmt.Println("value:", a+b)
}
