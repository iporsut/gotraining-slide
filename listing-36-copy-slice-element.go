package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	var b []int
	n := copy(b, a)
	fmt.Println(n)
	fmt.Println(b)

	b = make([]int, 2)
	n = copy(b, a)
	fmt.Println(n)
	fmt.Println(b)

	b = make([]int, 2)
	n = copy(b, a[2:])
	fmt.Println(n)
	fmt.Println(b)
}
