package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b [5]int
	fmt.Println(b)

	b = a
	fmt.Println(b)

	var c [6]int
	fmt.Println(c)

	c = a
	fmt.Println(c)
}
