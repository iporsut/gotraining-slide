package main

import "fmt"

func main() {
	prefix := "Hello"

	say := func(msg string) {
		fmt.Println(prefix, msg)
	}

	say("World")
}
