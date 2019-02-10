# go-examples

## Install

go get github.com/jamesacampbell/go-examples

## TL;DR;

cd to the directory and run: `go run main.go [instagram username]` and it will run each example

## What

Examples of libraries I use in golang frequently as a way for me to remember and hopefully useful for others making their way. Go get 'em. ;)

## 1 REDIS

The code is in redis1 folder, but just run main.go from root and it will run this along with the other examples. _You must have redis installed locally for this to work._

It should print `PONG` to show it is connected.

## 2 HTTP SERVER

The code is in the httpserver2 folder, it also runs from the root main.go file just fine and server Hello World to 127.0.0.1:1337 because you are a N00B. ;)

## 3 HEADLESS CHROME SCRAPER

The code is in the headless3 folder, it runs from main.go and asks you for an instagram profile to load. Instagram notoriously doesn't load if javascript is not available, thus a good test for headless chrome. It takes a instagram-profile.png to the root folder. *It only works in OSX / Mac OS* and if you have [Google Chrome](https://www.google.com/chrome/) installed already.

## 4 ENTITY EXTRACTOR USING PROSE

[Prose]("gopkg.in/jdkato/prose.v2") is pretty neat. It pulls out People, Places, and other things right out of the box. I tested it on English text and it seems to do pretty well. This example takes the text scraped from number 3 and throws it through Prose.
