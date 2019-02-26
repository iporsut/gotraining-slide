package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b []int
	b = a[1:3]
	fmt.Println(b, "len:", len(b), "cap:", cap(b))

	b = a[1:]
	fmt.Println(b, "len:", len(b), "cap:", cap(b))

	b = a[:3]
	fmt.Println(b, "len:", len(b), "cap:", cap(b))

	b = a[:]
	fmt.Println(b, "len:", len(b), "cap:", cap(b))
}
