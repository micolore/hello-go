package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	//初始化参数d 为间隔时间
	timer1 := time.NewTimer(2 * time.Second)
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1 ", time.Now())
		}

	}(ticker1)
	go func(t *time.Timer) {
		wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now())
			t.Reset(2 * time.Second)
		}
	}(timer1)

	wg.Wait()
}
