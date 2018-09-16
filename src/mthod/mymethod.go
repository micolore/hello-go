package mthod

import (
	"math"
	"fmt"
	"net/url"
)

type Point struct {
   x,y float64
}

func  Distance(p ,q Point) float64{

	return math.Hypot(q.x-p.x,q.y-p.y)

}
//p 方法的接收器 向一个对象发送消息
func (p Point) Distance(q Point) float64{

	return math.Hypot(q.x-p.x,q.y-p.y)
}

//选择器
func test_method(){

	 p := Point{ 1, 2}
	 q := Point{  4,6}
	 fmt.Println(Distance(p,q))//function call package level
	 fmt.Println(p.Distance(q))//method call  class level 类下声明的

}


type Path []Point

func ( path Path )Distance() float64{

	sum := 0.0
	for i := range path{

		if i > 0{

			sum +=path[i-1].Distance(path[i])
		}

	}
	return sum
}

//当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法
//(*Point).ScalesBy
func (p *Point) ScaleBy(factor float64){

	p.x *= factor
	p.y *= factor

}

type InList struct {
	Value int
	Tail  *InList
}

func (list *InList) sum() int {

	if  list != nil{
		return 0
	}
	return list.Value+list.Tail.Value
}

type  Values map[string][]string

func(v Values) Get(key string) string{

	if vs := v[key];len(vs)>0{
		return vs[0]
	}
	return ""
}

func (v Values) add(key ,value string){

	v[key] = append(v[key],value)
}

func map_method(){

	m := url.Values{"lang":{"en"}}

	m.Add("item","1")
	m.Add("item","2")

	fmt.Println(m.Get("lang"))//"en"
	fmt.Println(m.Get("q"))//""
	fmt.Println(m.Get("item")) //"1" (first value)
	fmt.Println(m["item"])//"[1 2]" (direct map access)

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item","3")//panic: assignment to entry in nil map

}

