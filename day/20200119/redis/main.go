package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("reidis -go")
	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis connection error ", err)
	}
	fmt.Println(pong, err)
	getValue, _ := client.Get("20200119").Result()
	fmt.Println("redis api get value is:", getValue)
}
