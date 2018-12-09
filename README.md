# go-examples
go get github.com/jamesacampbell/go-examples

## What

Examples of libraries I use in golang frequently as a way for me to remember and hopefully useful for others making their way. Go get 'em. ;)

## 1 REDIS

The code is in redis1 folder, but just run main.go from root and it will run this along with the other examples. You must have redis installed locally for this to work.

It should print `PONG` to show it is connected.

## 2 HTTP SERVER

The code is in the httpserver2 folder, it also runs from the root main.go file just fine and server Hello World to 127.0.0.1:1337 because you are a N00B. ;)

## 3 HEADLESS CHROME SCRAPER

The code is in the headless3 folder, it runs from main.go and asks you for an instagram profile to load. Instagram notoriously doesn't load if javascript is not available, thus a good test for headless chrome. It takes a screenshot.png to the root folder. It only works in OSX / Mac OS and if you have [Google Chrome Canary](https://www.google.com/chrome/canary/) installed already.
