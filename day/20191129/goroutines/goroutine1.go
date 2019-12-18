package main

import (
	"fmt"
	"time"
)

// go  goroutines coroutines
// Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的
// Go 协程通过通道来通信；协程通过让出和恢复操作来通信
func main() {
	fmt.Println(" in main()")
	go longWait()
	go shortWait()
	fmt.Println(" about to sleep in main()")

	time.Sleep(10 * 1e9)
	fmt.Println(" at the of main()")

}

func longWait() {
	fmt.Println(" beginning longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println(" end of longWait()")
}

func shortWait() {
	fmt.Println(" beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println(" end of shortWait()")
}
