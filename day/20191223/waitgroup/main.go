package main

import (
	"fmt"
	"sync"
	"time"
)

func handle(data int) {

	fmt.Println("handle params: ", data)
	time.Sleep(time.Second * 6)
	fmt.Println(" in handle....")
}

func main() {
	fmt.Println("waitgroup start ....")
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(data int) {
			handle(data)
			wg.Done()
		}(i)
	}
	fmt.Println("now wait.....")
	wg.Wait()
	fmt.Println("waitgroup end....")

}
