package main

import "fmt"

func main() {
	str := "Hello, World"
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		fmt.Println(i, ": ", str[i])
	}

	str = "สวัสดี"
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		fmt.Println(i, ": ", str[i])
	}
}
