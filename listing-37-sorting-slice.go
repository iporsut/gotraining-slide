package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	sort.Ints(a)
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Reverse(sort.IntSlice(a))
	fmt.Println(a)
}
