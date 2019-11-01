// inteface
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello interface!")

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Ipone)
	phone.call()
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am  nokia i can call you")
}

type Ipone struct {
}

func (Ipone Ipone) call() {
	fmt.Println("i am  iphone i can call you")
}
