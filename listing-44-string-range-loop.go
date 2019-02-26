package main

import (
	"fmt"
)

func main() {
	str := "สวัสดี"

	for i, v := range str {
		fmt.Println(i, v)
	}

	for i, v := range str {
		fmt.Printf("%d: %c\n", i, v)
	}
}
