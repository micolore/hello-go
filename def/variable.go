package def

import (
	"fmt"
	"math/rand"
)


var  i,j,k int
var  b,f,s = true, 2.3 ,"four"

func VarDef(){

	//var f ,err =  os.Open("filepath")
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(s)
	fmt.Println(f)


	//简短变量声明 名字:= 表达式
	//简短变量声明语句中必须至少要声明一个新的变量
	//anim :=gif.GIF{LoopCount: "nframes"}
	freq := rand.Float32() * 3.0;
	t := 0
	fmt.Println(freq)
	fmt.Println(t)

	//i := 100
	//var boiling float64 = 100
	//var  names []string
	//var  err error
	//var p image.Point
	//
	//o,ds := 100,200

}
