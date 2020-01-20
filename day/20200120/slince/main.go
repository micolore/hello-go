package main

import "fmt"

func main() {

	fmt.Println("slince-go")

	var arr1 [6]int
	var s2 []int = arr1[2:5]

	fmt.Println("s2 len:", len(s2))

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(s2); i++ {
		fmt.Printf("slince at %d is %d\n", i, s2[i])
	}

	fmt.Printf("the length of arr1 is %d\n", len(arr1))
	fmt.Printf("the length of s2 is %d\n", len(s2))
	fmt.Printf("the capacity of s2 is %d\n\n", cap(s2))

	//get elements from slice
	s2 = s2[0:4]
	for i := 0; i < len(s2); i++ {
		fmt.Printf("the new slice at %d is %d\n", i, s2[i])
	}

	fmt.Printf("the length of s2 is %d\n", len(s2))
	fmt.Printf("the capacity of s2 is %d\n", cap(s2))

	var sumSlice = [5]int{0, 1, 2, 4}
	r := sum(sumSlice[:])

	fmt.Println("sum slice return value is ", r)

	byteGetIndexSlice := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	getIndex(byteGetIndexSlice)
	newSlice()
}

func sum(s []int) int {

	r := 0
	for i := 0; i < len(s); i++ {
		r += s[i]
	}
	return r
}

func getIndex(s []byte) {

	// index 1-4
	fmt.Println("slice 1:4 is", s[1:4])
	// index 0-2
	fmt.Println("slice :2 is ", s[:2])
	// index 2-length
	fmt.Println("slice 2: is ", s[2:])
	// all
	fmt.Println("slice : is ", s[:])

}

func newSlice() {
	// var slice1 []type = make([]type,len)
	s1 := make([]int, 50, 100)
	s2 := new([100]int)[0:20]
	fmt.Println(len(s1))
	fmt.Println(len(s2))

}
