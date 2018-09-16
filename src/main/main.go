package main

import  (

 		"fmt"
 		"myfunction"
	"myerror"
)

func main() {

	//def.Say()

	//def.Declare()

	//def.VarDef()

	//def.Guide()

	 //def.NewGo()

	 //def.Echo1()

	fmt.Println(myfunction.MyHypot(3,5))

	fmt.Printf("%T\n",myfunction.Add)
	fmt.Printf("%T\n",myfunction.Sub)
	fmt.Printf("%T\n",myfunction.First)
	fmt.Printf("%T\n",myfunction.Zero)

	myfunction.ManyResult()

	myerror.Err1("https://github.com/golang/net")


}


