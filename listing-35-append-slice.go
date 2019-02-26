package main

import "fmt"

func main() {
	var a []int
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 2)
	fmt.Println(len(a), cap(a))

	a = append(a, 3)
	fmt.Println(len(a), cap(a))

	a = append(a, 4, 5, 6)
	fmt.Println(len(a), cap(a))
}
