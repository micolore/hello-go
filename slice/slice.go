// slice 切片是对数组的抽象。
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello slice!")

	// var identifier []type
	// var slice1 []type = make([]type,len)
	// slice1 := make([]type ,len)
	// make([]T,length,capacity)
	//s := []int{1, 2, 3}

	var number = make([]int, 3, 5)

	printSlice(number)

	numbers := []int{23, 6, 3, 15, 0, 41}

	printSlice(numbers)

	split_slice(numbers)

	printSlice(numbers)

	number1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(number1, numbers)

	printSlice(number1)

}

func printSlice(x []int) {

	fmt.Printf("len=%d ,cap=%d slice=%v\n", len(x), cap(x), x)
}

func split_slice(n []int) {

	fmt.Println("numbers[1:4]", n[1:4])

	fmt.Println("numbers[:3]", n[:3])

	fmt.Println("numbers[4:]", n[4:])

	n = append(n, 1)

	printSlice(n)

}
