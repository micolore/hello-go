package main

import (
	"fmt"
	"time"
)

// channel 通道内没有消费者，线程都是阻塞的（死锁）

//同步通道
//buf := 100
//ch1 := make(chan string, buf)
//缓冲容量和类型无关
//如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功
//若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable）。要在首要位置使用无缓冲通道来设计算法，只在不确定的情况下使用缓冲
func main() {

	//ch := make(chan int)
	//go sendData(ch)
	//go getData(ch)
	ch1 := make(chan int)
	go sum(ch1)
	time.Sleep(1e9 * 2)

	go printSum(ch1)
	fmt.Println("send ok ..")
}

func sendData(ch chan int) {

	for i := 0; i < 1; i++ {
		ch <- i
	}
}

func getData(ch chan int) {
	for {
		value := <-ch
		fmt.Println(value)
	}
}

func sum(ch chan int) {
	ch <- 3
	var value int = 2
	ch <- value

}
func printSum(ch chan int) {

	for {
		value := <-ch
		fmt.Println(value)
		fmt.Println("printSum")
	}

}
