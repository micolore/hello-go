package main

import "fmt"

func Afunction(routineCtl chan int, feedback chan string) {
	temp := <-routineCtl
	fmt.Println("Afunction start...", temp)
	defer func() {
		feedback <- "finsh"
		fmt.Println("Afunction end...", temp)
	}()
	//dosomething
}

func main() {
	fmt.Println("channel10 start...")
	var (
		routineCtl chan int    = make(chan int, 20)
		feedback   chan string = make(chan string, 10000)
		msg        string
		allwork    int
		finshed    int
	)
	for i := 0; i < 1000; i++ {
		routineCtl <- i
		allwork++
		go Afunction(routineCtl, feedback)
	}

	for {
		msg = <-feedback
		if msg == "finsh" {
			finshed++
		}
		if finshed == allwork {
			break
		}
	}
	fmt.Println("channel10 end...")
}
