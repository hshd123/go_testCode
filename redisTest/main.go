package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var myRedisClient *redis.Client

func redisInit() *redis.Client {
	fmt.Println("regist redis db ...", time.Now().String())
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		PoolSize:     5,
		PoolTimeout:  time.Second * 10,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		DialTimeout:  time.Second * 10,
	})
	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("redis connect state ", ping)
	s1 := strings.Compare(ping, "PONG")
	if s1 == 0 {
		fmt.Println("connect succ ")
	}
	return client
}

func main() {
	fmt.Println("hello world!")
	myRedisClient = redisInit()
	var s string = ""
	for i := 0; i < 256; i++ {
		s = s + "-" + string(i)
	}
	myRedisClient.Set("s1", s, 0)
	rec := myRedisClient.Get("s1")
	if rec.Err() != nil {
		myRedisClient.Set("s1", s, 0)
	}

	fmt.Println("rsl-- ", myRedisClient.Get("s1").Val())

}
