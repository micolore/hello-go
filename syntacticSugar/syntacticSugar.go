package main

import "fmt"

//语法糖 用于编程语言中的某种类型的语法，如:= ...

func main() {

	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j=%d\n", i, j)
}

// 1、 := 多变量声明会重新赋值
// 2、不能用于函数外部
//
