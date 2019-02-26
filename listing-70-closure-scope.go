package main

import "fmt"

func makeSayFunc(prefix string) func(string) {
	return func(msg string) {
		fmt.Println(prefix, msg)
	}
}

func main() {
	say := makeSayFunc("Hello")
	say("World")
}
