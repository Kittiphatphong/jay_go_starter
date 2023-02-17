package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_starter/config"
	"go_starter/logs"
)

func InitRedis() (*redis.Client, error) {

	redisHost := config.GetEnv("redis.host", "localhost")
	redisPORT := config.GetEnv("redis.port", "6379")
	fmt.Println("CONNECTING_TO_REDIS_SERVER: ", redisHost, ":", redisPORT)
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPORT,
		Password: "",
		DB:       0,
		//Addr: "host.docker.internal:6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logs.Error(err)
		return nil, errors.New("CANT_CONNECT_TO_REDIS")
	}
	fmt.Println("CONNECTED_TO_REDIS")
	return client, nil
}
