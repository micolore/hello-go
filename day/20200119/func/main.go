package main

import "fmt"

type Common struct {
}

func add(i int) int {
	return i + 1
}

func (c Common) add(i int) int {
	return i + 1
}

func addN() func(p int) int {
	//var x int = 0
	//声明周期
	x := 0
	return func(p int) int {
		fmt.Println("addN none func name x value is:", x)
		x += p
		return x
	}
}

func sub(i int) int {
	return i - 1
}

func paramFunc(i int, f func(p int) int) int {
	return f(i)
}
func main() {
	fmt.Println("func go")
	var p int = 88
	f := add(p)
	fmt.Printf("first value is:%d param is:%d\n", f, p)
	// 生命周期
	nf1, nf2 := addN(), addN()

	for i := 0; i < 3; i++ {
		nfValue := nf1(i)
		nfValue2 := nf2(-2 * i)
		fmt.Printf("none func name return value1 is:%d value2 is:%d\n", nfValue, nfValue2)
	}

	common := Common{}
	commonValue := common.add(12)
	fmt.Println("common return value is:", commonValue)

	// 方法表达式
	varAdd := add
	varAddValue := varAdd(13)
	fmt.Println("var add value:", varAddValue)

	paramFuncValue := 11
	paramFunAddValue := paramFunc(paramFuncValue, add)
	paramFunSubValue := paramFunc(paramFuncValue, sub)
	fmt.Printf("paramFunAddValue params value:%d return value:%d \n", paramFuncValue, paramFunAddValue)
	fmt.Printf("paramFunSubValue params value:%d return value:%d \n", paramFuncValue, paramFunSubValue)

	dc := deferClose()
	fmt.Println("dc value is", dc)

	forNoneFunc()
}

func deferClose() (r int) {
	//闭包对捕获的外部变量并不是传值的方式访问，而是以引用的方式访问
	fmt.Println("deferClose start r value is :", r)
	defer func() {
		fmt.Println("deferClose none name function value is:", r)
		r++

	}()
	fmt.Println("deferClose end r value is :", r)
	return 23
}

func forNoneFunc() {

	for i := 0; i < 5; i++ {
		i := i
		defer func() { println(i) }()
	}

	for i := 0; i < 5; i++ {
		defer func(p int) {
			fmt.Println(p)
		}(i)
	}
}
