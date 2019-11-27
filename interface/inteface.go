// inteface
package main

import (
	"fmt"
)

/**
 *
 * 1，接口用来 说明对象的行为
 * 2，接口可以嵌套接口
 */
func main() {
	fmt.Println("Hello interface!")

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone.recive()

	//动态的实现（多态）
	phone = new(Ipone)
	phone.call()
	phone.recive()
}

type Phone interface {
	call()
	recive()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am  nokia i can call you")
}
func (nokiaPhone NokiaPhone) recive() {
	fmt.Println("i am  nokiaPhone i can receive you")
}

type Ipone struct {
}

// 这行的作用是为了给iphone写call的实现
// 其实就是给iphone加了一个call方法
func (Ipone Ipone) call() {
	fmt.Println("i am  iphone i can call you")
}

func (Ipone Ipone) recive() {
	fmt.Println("i am  iphone i can receive you")
}
