package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b []int
	fmt.Println(b)

	b = a
	fmt.Println(b)

	b[0] = 20
	fmt.Println(b)
	fmt.Println(a)
}
