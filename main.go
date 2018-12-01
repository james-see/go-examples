package main

import (
	http2 "github.com/jamesacampbell/go-examples/http2"
	redis1 "github.com/jamesacampbell/go-examples/redis1"
)

func main() {
	redis1.Redis1()
	http2.HTTPExample()
}
