package main

import "fmt"

func main() {
	a := 10
	b := 20

	var pa *int // pa is nil
	pa = &a     // pa point to a
	fmt.Println(pa)
	pb := &b // pb point to b
	fmt.Println(pb)

	fmt.Println("a:", a, "*pa:", *pa)
	*pa = 50
	fmt.Println("a:", a, "*pa:", *pa)

	fmt.Println("b: ", b, "*pb", *pb)
	*pb *= 2
	fmt.Println("b: ", b, "*pb", *pb)
}
