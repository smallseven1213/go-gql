package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var REDIS *redis.Client

func RedisInit() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Using this function to get a connection, you can create your connection pool here.
func GetRedis() *redis.Client {
	return REDIS
}
