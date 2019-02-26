package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func do(op func(int, int) int, x int, y int) {
	fmt.Println(op(x, y))
}
func main() {
	do(add, 10, 20)
	do(sub, 50, 10)
}
