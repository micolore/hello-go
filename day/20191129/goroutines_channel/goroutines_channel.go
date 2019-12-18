package main

// channel  类型队列（先进先出）
// var ch1 := make(chan,string)
// ch <- int1  向ch通道发送int1
// int2 = <- ch  从ch通道获取数据

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {

	ch <- "nba"
	ch <- "cba"
	ch <- "wcba"
}
func getData(ch chan string) {

	var input string

	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}
