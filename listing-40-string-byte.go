package main

import "fmt"

func main() {
	str := "สวัสดี"
	b := []byte(str)
	fmt.Println(b)

	b = []byte{224, 184, 170, 224, 184, 167, 224, 184, 177, 224, 184, 170, 224, 184, 148, 224, 184, 181}
	str = string(b)
	fmt.Println(str)
}
