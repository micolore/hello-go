// error
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello error!")

	value, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func sqrt(f float32) (float32, error) {

	if f < 0 {
		return 0, errors.New("math: sqrt roor of nagative number")
	}
	return f, nil
}
