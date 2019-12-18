package main

import "fmt"

func main() {
	fmt.Println("test-vim")
	fmt.Print("good")
	var name = "li"
	fmt.Println(name)
	d1 := dog{"wangcai", 1}
	fmt.Println(d1.name)
	fmt.Println(d1.age)
}

type dog struct {
	name string
	age  uint32
}
