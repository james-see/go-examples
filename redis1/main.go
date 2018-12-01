package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Connect to local redis instance and return Client object
func redisCon() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return r
}

func main() {
	r := redisCon() // returned client object
	// test ping pong command
	pong, _ := redisdb.Ping().Result()
	fmt.Println(pong) // should print PONG to console

}
