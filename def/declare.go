package def

import "fmt"

//包一级的变量声明
const  boilindF  = 213.0

func Declare(){
	var e,g float32
	var f = boilindF
	var c = (f - 32)*5/9;
	fmt.Println("boiling print = %g F or %g C\n",f,c)

	fmt.Println("boiling print = %g F or %g C\n",e,g)
}