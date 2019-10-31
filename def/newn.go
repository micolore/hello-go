package def

import "fmt"

func NewGo(){

	p:= new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	fmt.Println(newInt() ==  newInt2())
}

func newInt() *int{
	return new(int)
}
func newInt2() *int{
	var dummy int
	return &dummy
}

// x 并不会因为函数退出 就会回收
var global *int

func fout(){

	var x int
	x = 1
	global = &x;
}
// 函数退出就会立即回收
func fin(){
	var  y = new(int)

	*y = 1
}