package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func main() {
	var operator func(int, int) int
	fmt.Println(operator == nil) // zero value is nil

	operator = add
	fmt.Println(operator(20, 30))

	operator = sub
	fmt.Println(operator(30, 20))
}
