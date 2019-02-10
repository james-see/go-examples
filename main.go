package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	headless3 "github.com/jamesacampbell/go-examples/headless3"
	headliner4 "github.com/jamesacampbell/go-examples/headliner4"
	http2 "github.com/jamesacampbell/go-examples/httpserver2"
	redis1 "github.com/jamesacampbell/go-examples/redis1"
)

func main() {
	// example 1
	fmt.Println("Example 1 connects to Redis at local port 6379 and ping / pong.")
	redis1.Redis1()
	var username string
	if len(os.Args) == 2 {
		username = os.Args[1]
	} else {
		username = "defaultuser"
	}
	// example 2
	fmt.Println("Example 2 should be running at http://127.0.0.1:1337")
	fmt.Println("Runs for 10 seconds so you can refresh browser and test to see Hello World.")

	go http2.HTTPExample()
	time.Sleep(10 * time.Second)
	// example 3
	fmt.Println("Example 3 screenshot's instagram account saved to instagram-profile.png for example 3.")
	instatext := headless3.Headless4(username)
	var words []string
	l := 0
	for s := instatext; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}
	var fixedstring string
	for i := 0; i < len(words); i++ {
		//fmt.Printf("%s finished as number %d\n", words[i], i+1)
		fixedstring = fixedstring + " " + words[i]
	}
	fmt.Println("Here is the text extracted from Instagram account:")
	fmt.Println(fixedstring)
	// example 4
	fmt.Println("Example 4 runs entity extractor on the Instagram text.")
	headliner4.GetEntities(fixedstring)
	// More coming soon
	os.Exit(0)
}
