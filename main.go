package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const redisUrl = "localhost:6379"

func main() {
	redisOptions := &redis.Options{
		Addr:         redisUrl,
		MaxRetries:   3,
		MinIdleConns: 50,
		MaxConnAge:   30 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		DB:           0, // use default DB
	}

	client := redis.NewClient(redisOptions)
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}
	defer client.Close()
	ctx := context.Background()
	err = redis.SAdd(ctx, "user:id:name", "user:1:John", "user:2:Jack").Err()
	if err != nil {
		fmt.Println("redis addd err:", err)
	}
	redisValExists := redis.SIsMember(ctx, "user:id:name", "user:1:John").Val()
	fmt.Println(redisValExists)
}
