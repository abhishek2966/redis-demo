package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

const redisUrl = "localhost:6379"

func main() {
	rdb := newRedisConnection()
	defer rdb.Close()

	ctx := context.Background()

	// working with redis Strings
	//sets "reqCount:example.com" = 1 to be expired in an hour
	err = rdb.Set(ctx, "reqCount:example.com", 1, time.Hour).Err()
	if err != nil {
		fmt.Println("redis addd err:", err)
	}
	//increments value of "reqCount:example.com" by 1

	err = rdb.SAdd(ctx, "user:id:name", "user:1:John", "user:2:Jack").Err()
	if err != nil {
		fmt.Println("redis addd err:", err)
	}
	redisValExists := rdb.SIsMember(ctx, "user:id:name", "user:1:John").Val()
	fmt.Println(redisValExists)
}

func newRedisConnection() *redis.Client {
	redisOptions := &redis.Options{
		Addr: redisUrl,
		DB:   0, // use default DB
	}
	return redis.NewClient(redisOptions)
}
