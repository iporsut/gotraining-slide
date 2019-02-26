package main

import "fmt"

func main() {
	defer fmt.Println("Done")
	defer fmt.Println("Before Done")

	fmt.Println("Hello, World")
}
