// func
package main

//引入指定的包
import "fmt"

/*
 * go同一个目录下面只允许有一个main
 */
func main() {
	//go 声明变量
	var a int = 10
	var b int = 20
	//函数调用
	var result int = say(a, b)
	//类似c的格式化输出
	fmt.Printf("result:%d\n", result)

	//调用方法
	var area_count, area_size = say_more(10, 20, 30)

	fmt.Printf("area_count: %d\narea_size: %d\n", area_count, area_size)
}

/*
 * go声明方法
 */
func say(count int, size int) int {

	var result int
	result = count * size
	return result
}

/*
 * 方法返回多个返回值
 */
func say_more(count int, size int, area int) (int, int) {

	var r_area_count int = count * area
	var r_area_size int = size * area
	return r_area_count, r_area_size
}
