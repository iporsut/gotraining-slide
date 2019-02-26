package main

import "fmt"

func main() {
	str := "สวัสดี"
	r := []rune(str)
	fmt.Println(r)

	r = []rune{3626, 3623, 3633, 3626, 3604, 3637}
	str = string(r)
	fmt.Println(str)
}
