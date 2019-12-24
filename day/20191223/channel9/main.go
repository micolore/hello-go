package main

import "fmt"

func Afunction(ch chan int) {
	fmt.Println("channel9 Afunction start ...")
	ch <- 1
}
func main() {
	fmt.Println("channel9 start...")
	var (
		ch        chan int = make(chan int, 20)
		dutycount int      = 500
	)
	for i := 0; i < dutycount; i++ {
		go Afunction(ch)
	}
	for i := 0; i < dutycount; i++ {
		<-ch
	}
	fmt.Println("channel9 end...")
}
