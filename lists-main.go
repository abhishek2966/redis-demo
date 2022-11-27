package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

const redisUrl = "localhost:6379"

// Redis lists are linked list
func main() {
	rdb := newRedisConnection()
	defer rdb.Close()

	ctx := context.Background()

	//left push into list
	rdb.LPush(ctx, "taskqueue", "task1", "task2", "task3") // ["task3" "task2" "task1"]
	rdb.LPush(ctx, "taskqueue", "task2")                   // ["task2" "task3" "task2" "task1"]

	//right push into list
	rdb.RPush(ctx, "taskqueue", "task4") // ["task2" "task3" "task2" "task1" "task4"]

	// LPUSH & RPOP makes a queue
	// LPUSH & LPOP makes a stack

	//left pop from list
	fmt.Println(rdb.LPop(ctx, "taskqueue").Val()) // "task2"
	//right pop from list
	fmt.Println(rdb.RPop(ctx, "taskqueue").Val()) // "task4"

}

func newRedisConnection() *redis.Client {
	redisOptions := &redis.Options{
		Addr: redisUrl,
		DB:   0, // use default DB
	}
	return redis.NewClient(redisOptions)
}
