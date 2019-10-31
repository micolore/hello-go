// pointer
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Pointer!")

	var a int = 10
	fmt.Printf("variable address %x\n", &a)

	//var ip *int
	//var fp *float32
	use_pointer()

	nil_pointer()

}

func use_pointer() {

	var a int = 10
	var ip *int
	ip = &a
	fmt.Printf("a variable address: %x\n", &a)

	fmt.Printf("ip variable address %x\n", ip)

	fmt.Printf("ip variable value is :%d\n", *ip)
}

func nil_pointer() {

	var long *int

	fmt.Printf("long value is :%x\n", long)

	//if(long!=nil)
	//if(long == nil)

}
