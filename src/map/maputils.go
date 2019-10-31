// maputils
package main

import (
	"fmt"
	"math"
)

var PI float64 = 3.14159265
var EARTH_RADIUS float64 = 6378137
var RAD float64 = math.Pi / 180.0

func main() {
	fmt.Println("Hello map!")
	var rad float64 = rad(10222.021)
	fmt.Printf("rad %d\n\n", rad)
	var distance float64 = distance(120.179128, 30.285243, 121.384151, 31.222716)
	fmt.Println("distince:", distance)

}

/*
 * 根据经纬度计算距离，返回单位km
 */

func distance(lon_one float64, lat_one float64, lon_two float64, lat_two float64) float64 {

	var radLatOne float64 = rad(lat_one)
	var radLatTwo float64 = rad(lat_two)

	var radLonOne float64 = rad(lon_one)
	var radLonTwo float64 = rad(lon_two)
	var radLatA float64 = radLatOne - radLatTwo
	var radLonA float64 = radLonOne - radLonTwo

	var result float64 = 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(radLatA/2), 2)+math.Cos(radLatOne)*math.Cos(radLatTwo)*math.Pow(math.Sin(radLonA/2), 2)))
	result = result * EARTH_RADIUS
	return result / 1000
}

/*
 * 获取指定范围的坐标(最大、最小经纬度)
 */
func around(latitude float64, longitude float64, raidus float64) (float64, float64, float64, float64) {

	var degree float64 = (24901 * 1609) / 360.0
	var raidusMile float64 = raidus

	var dpmLat float64 = 1 / degree
	var radiusLat float64 = dpmLat * raidusMile
	var minLat float64 = latitude - radiusLat
	var maxLat float64 = latitude + radiusLat

	var mpdLng float64 = degree * math.Cos(latitude*(PI/180))
	var dpmLng float64 = 1 / mpdLng
	var radiusLng float64 = dpmLng * raidusMile
	var minLng float64 = longitude - radiusLng
	var maxLng float64 = longitude + radiusLng

	return minLat, minLng, maxLat, maxLng
}

func rad(rad float64) float64 {

	return rad * math.Pi / 180.0
}
