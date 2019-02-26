package main

import "fmt"

func main() {
	var truth bool
	fmt.Println(truth)
	// operator การเปรียบเทียบ ==, <=, >= , <, >, !=
	truth = 10 == 10
	fmt.Println(truth)

	// ! เป็นนิเสธ กลับค่าจาก true เป็น false และ false เป็น true
	fmt.Println(!truth)
	fmt.Println(!true, !false)
}
