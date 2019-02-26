package main

import "fmt"

const prefix = "Hello, "
const maxBuffer = 2048

func sayHello(s string) {
	fmt.Println(prefix + s)
}

func main() {
	sayHello("World")
}
