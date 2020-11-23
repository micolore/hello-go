package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

const redisAddr = "localhost:6379"

func init() {
	initSingle()
}

// redis
func initSingle() {
	op := &redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	}
	rdb := redis.NewClient(op)
	err := rdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("set error!")
	}
	fmt.Printf("%s init success!\n", "redis")
}

func initCluster() {

}

func main() {
}
