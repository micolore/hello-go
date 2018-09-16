package def



//每个源文件都是以包的声明语句开始，用来指名包的名字。当包被导入的时候，包内的成员将通过类似
//tempconv.CToF的形式访问。而包级别的名字，例如在一个文件声明的类型和常量，在同一个包的其他源
//文件也是可以直接访问的，就好像所有代码都在一个文件一样。要注意的是tempconv.go源文件导入了
//fmt包，但是conv.go源文件并没有，因为这个源文件中的代码并没有用到fmt包

func Ctof(c Celius) Fahrenheit{
	return Fahrenheit(c*9/5+32)
}

func package_init(){

	//var a = b+c; a 3
	//var b = f()  b 2
	//var c =1  c 1
}

func init(){

	// i will init code

	//初始化工作是自下而上进行的

}
