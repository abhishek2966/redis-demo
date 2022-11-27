package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

const redisUrl = "localhost:6379"

func main() {
	redisOptions := &redis.Options{
		Addr: redisUrl,
		DB:   0, // use default DB
	}

	ctx := context.Background()

	rdb := redis.NewClient(redisOptions)
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}
	defer rdb.Close()

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

	//sets string user1 value user1@email.com
	rdb.Set(ctx, "user1", "user1@email.com")

	fmt.Println(rdb.Get(ctx, "reqCount:example.com").Result())
	fmt.Println(rdb.Get(ctx, "reqCount:google.com").Result())

	err = rdb.SAdd(ctx, "user:id:name", "user:1:John", "user:2:Jack").Err()
	if err != nil {
		fmt.Println("redis addd err:", err)
	}
	redisValExists := rdb.SIsMember(ctx, "user:id:name", "user:1:John").Val()
	fmt.Println(redisValExists)
}
