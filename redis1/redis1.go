package redis1

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisCon : to local redis instance and return Client object
func RedisCon() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return r
}

// Redis1 : Connect to redis
func Redis1() {
	r := RedisCon() // returned client object
	// test ping pong command
	pong, _ := r.Ping().Result()
	fmt.Println(pong) // should print PONG to console

}
