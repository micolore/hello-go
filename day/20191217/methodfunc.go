package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		i := i
		defer func() {
			fmt.Println("decrlere:", i)
		}()
	}
	for x := 0; x < 3; x++ {
		defer func(i int) { fmt.Println("params:", x) }(x)
	}
}
