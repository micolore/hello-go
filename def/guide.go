package def

import "fmt"

func Guide(){

	x := 1
	//&x表达式（取x变量的内存地址）
	p := &x
	//将产生一个指向该整数变量的指针，指针对应的数据类型是*int
	//指针被称之为 "指向int类型的指针"
	fmt.Println(*p)
	//一般*p表达式读取指针指向的变量的值，这里为int类型的值，同时因为*p对应一个
	//变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值
	*p  = 2
	fmt.Println(x)
	//任何类型的指针的零值都是nil

	var f , y , z  int
	fmt.Println( &f == &f ,&f == &z ,&y == nil)

	var xx = foo()
	fmt.Println(xx)
	fmt.Println(foo() == foo())

	iv := 1;
	var ii = incr(&iv)
	fmt.Println(ii)
	fmt.Println(incr(&iv))
}

func foo() *int{
	v := 1
	return &v
}

func incr(pp *int) int{
	*pp++
	return *pp
}