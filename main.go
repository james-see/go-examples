package main

import (
	"fmt"

	http2 "github.com/jamesacampbell/go-examples/httpserver2"
	redis1 "github.com/jamesacampbell/go-examples/redis1"
)

func main() {
	fmt.Println("Example 1 connects to Redis at local port 6379 and ping / pong.")
	redis1.Redis1()
	fmt.Println("Example 2 should be running at http://127.0.0.1:1337")
	fmt.Println("CTRL + C to exit")
	http2.HTTPExample()
	// More coming soon
}
