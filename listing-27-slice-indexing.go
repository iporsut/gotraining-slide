package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a[3]) // cannot compile

	s := []int{1, 2, 3}
	fmt.Println(s[3]) // panic
}
