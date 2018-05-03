package main

import "github.com/AlfredDobradi/kit/logging"

func main() {
	l := logging.New(true)
	l.Debug("hai")
}
