package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b [5]int
	fmt.Println(b)

	b = a
	fmt.Println(b)

	b[0] = 20
	fmt.Println(b)

	fmt.Println(a)
}
