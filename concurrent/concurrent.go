package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	channels := make([]chan int, 1)

	for i := 0; i < 1; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}
	for i, ch := range channels {
		<-ch
		fmt.Println("routine", i, "quit!")
	}

	WaitGroup()
}

// Process ...
func Process(ch chan int) {

	time.Sleep(time.Second * 1)

	//write a  element means it's end
	ch <- 1
}

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("G1 finished!")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("G2 finished!")
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("all goroutines is fishnished!")
}

type EmptyCtx int

func (e *EmptyCtx) Deadline(deadline time.Time, ok bool) {
	return
}

func (e *EmptyCtx) Done() <-chan struct{} {
	return nil
}

func (e *EmptyCtx) Err() error {
	return nil
}

func (e *EmptyCtx) Value(key interface{}) interface{} {
	return nil
}
