package def

import "fmt"

//type  类型名字 底层类型
//通常出现在包一级 如果包类型的名字首字母大写 那外部包也可以访问

type Celius float64
type Fahrenheit float64

//对于每一个类型T，都有一个对应的类型转换操作T(x)
//只有当两个类型的底层基础类型相同时，才允许这种转型操作
const(
	AbsoluteZeroC Celius = -273.5
	FreezingC  Celius = 0
	BoilingC   Celius = 100
)

func CToF(c Celius) Fahrenheit {

	return Fahrenheit(c*9/5+32)
}

func CToC(f Fahrenheit) Celius{

	return Celius((f-32) * 5 / 9)

}

func cast(){
	var c Celius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println( c == f) compile error type mismatch
	fmt.Println( c == Celius(f))

}

//Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个叫名叫
//String的方法，该方法返回该类型对象c带着°C温度单位的字符串
func (c Celius) String() string {

	return fmt.Sprintln("%g C",c)

}


