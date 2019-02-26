package main

import "fmt"

func main() {
	str := "สวัสดี"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i]) // string index is byte index
	}
}
