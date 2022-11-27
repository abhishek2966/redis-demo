package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

const redisUrl = "localhost:6379"

// Redis Set is an unordered collection of unique strings
func main() {
	rdb := newRedisConnection()
	defer rdb.Close()

	ctx := context.Background()

	// add few elements to the set named userSet
	rdb.SAdd(ctx, "userSet", "user:1:John", "user:2:Jack")
	rdb.SAdd(ctx, "userSet", "user:3:Ryan", "user:4:Tim", "user:1:John")

	//check if an element is present in userSet
	redisValExists := rdb.SIsMember(ctx, "userSet", "user:1:John").Val()
	fmt.Println(redisValExists) //true
	//removes few elements from usrset
	rdb.SRem(ctx, "userSet", "user:2:Jack", "user:4:Tim")

	//cardinality
	fmt.Println(rdb.SCard(ctx, "userSet")) // 2
}

func newRedisConnection() *redis.Client {
	redisOptions := &redis.Options{
		Addr: redisUrl,
		DB:   0, // use default DB
	}
	return redis.NewClient(redisOptions)
}
