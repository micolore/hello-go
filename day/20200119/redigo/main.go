package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("redis-go")

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connection error ...", err)
	}
	fmt.Println("redis connecion  success!")
	defer c.Close()

	setKeyStr := "20200119"
	c.Do("set", setKeyStr, 2020)
	n, _ := redis.Int(c.Do("get", setKeyStr))
	fmt.Println("set redis int get value:", n)
}
