package main

import (
	"fmt"
	"time"
)

func main() {
	firstTimer()
	timeAfter()
	firstTicker()
	time.Sleep(10 * time.Second)

}

func firstTimer() {

	fmt.Printf("start at %s\n", time.Now().String())
	t := time.NewTimer(time.Duration(5 * time.Second))

	select {
	case <-t.C:
		fmt.Printf("first timer execute at %s\n", time.Now().String())
	}
}
func timeAfter() {
	time.AfterFunc(time.Second*4, func() {
		fmt.Printf("timeAfter execute %s\n", time.Now().String())
	})
}

// ticker定时器
func firstTicker() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	for range t.C {
		fmt.Printf("fistTicker execute %s\n", time.Now().String())
	}
}
