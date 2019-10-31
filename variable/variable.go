// variable
package main

import (
	"fmt"
)

//declare global variable
var g int
var gt int = 20

func main() {
	fmt.Println("Hello World!")

	//declare variable
	var a, b, c int

	//init variable
	a = 10
	b = 20
	c = a + b
	g = a + b + c
	fmt.Printf("result: a=%d ,b= %d ,c=%d ,g=%d\n", a, b, c, g)

	fmt.Printf("gt=%d\n", gt)

	var sumValue = sum(a, b)
	fmt.Printf("sum value %d\n", sumValue)
}

func sum(a, b int) int {

	return a + b
}
