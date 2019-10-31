package mthod

import (
	"image/color"
	"fmt"
)

type Point struct {
	x,y float64
}

type ColredPoint struct {

	Point
	Color color.RGBA
}

func m1(){

	var cp ColredPoint
	cp.x=1
	fmt.Println(cp.Point.x)
	cp.Point.y =2
	fmt.Println(cp.y)

	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}

	var p = ColredPoint{Point{1,1},red}
	var q = ColredPoint{Point{5,4},blue}
	

}