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
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}

	// working with redis Strings
	//sets "reqCount:example.com" = 1 to be expired in an hour
	err = rdb.Set(ctx, "reqCount:example.com", 1, time.Hour).Err()
	if err != nil {
		fmt.Println("redis addd err:", err)
	}
	//increments value of "reqCount:example.com" by 1
	rdb.Incr(ctx, "reqCount:example.com")

	//increments value of "reqCount:example.com" by 4
	rdb.IncrBy(ctx, "reqCount:example.com", 4)

	//sets string user1 value user1@email.com with no expiration
	rdb.Set(ctx, "user1", "user1@email.com", 0)

	fmt.Println(rdb.Get(ctx, "reqCount:example.com").Result()) // 6 nil

	val, _ := rdb.Get(ctx, "reqCount:google.com").Val() // val = ""
	fmt.Println("reqCount:google.com", val)

	// if key doesn't exist
	if rdb.Get(ctx, "reqCount:google.com").Err() == redis.Nil {
		fmt.Println("reqCount:google.com key doen't exist")
	}

}

func newRedisConnection() *redis.Client {
	redisOptions := &redis.Options{
		Addr: redisUrl,
		DB:   0, // use default DB
	}
	return redis.NewClient(redisOptions)
}
