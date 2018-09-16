package myfunction

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
)

func square(n int) int {return n*n}

func negative( n int) int{return -n}

func product(m,n int) int {return m*m}

func func_type(){

	f := square
	fmt.Println(f(2))

	s := negative
	fmt.Println(s(3))
	fmt.Printf("%T\n",f)

	//s := product compile error: can't assign func(int, int) int to func(int) int
	//var f func(int) int
	//f(2)

	f2()
	f3()
}
func f2(){
	var f func(int) int
	if f != nil{

		f(2)
	}
}

func  add1(r rune) rune{
	return r+1
}

func f3(){

	fmt.Println(strings.Map(add1,"111"))
	fmt.Println(strings.Map(add1,"222"))
	fmt.Println(strings.Map(add1,"333"))

}

func  forEachNode(n *html.Node ,pre,post func(n *html.Node)){

	if pre != nil{
		pre(n)
	}

	for c := n.FirstChild; c!=nil ;c=c.NextSibling{

		forEachNode(c,pre,post)

	}

	if post != nil{
		post(n)
	}

}


var  depth int



func  startElement( n *html.Node){

	if n.Type == html.ElementNode {

		fmt.Printf("%*<%s\n",depth*2,"",n.data)
		depth++
	}

}
//%*s中的*会在字符串之前填充一些空格
func endElement(n *html.Node)  {

	if n.Type != html.ElementNode{
		depth--
		fmt.Printf("%*s</%s>\n",depth*2,"",n.Data)
	}
}

//匿名函数
func noName(){

	strings.Map(func(r rune) rune {return r+1},"wow key")

}

//squares返回一个匿名函数
// 该匿名函数每次被调用时都会返回下一个数的平方
//函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内部函数可以
//访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就是函数值属于
//引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。
func squares() func() int{

	var x int

	return func() int{
		x++
		return x *x
	}

}

//defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机
//制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后



//当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成线程，在
//第8章会详细介绍）中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。日志信息包括
//panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信息。对于每个goroutine，日志
//信息中都会有与之相对的，发生panic时的函数调用堆栈跟踪信息。通常，我们不需要再次运行程序去定
//位问题，日志信息已经提供了足够的诊断依据。因此，在我们填写问题报告时，一般会将panic异常和日志信息一并记录


//当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成线程，在
//第8章会详细介绍）中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。日志信息包括
//panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信息。对于每个goroutine，日志
//信息中都会有与之相对的，发生panic时的函数调用堆栈跟踪信息。通常，我们不需要再次运行程序去定
//位问题，日志信息已经提供了足够的诊断依据。因此，在我们填写问题报告时，一般会将panic异常和日志信息一并记录
