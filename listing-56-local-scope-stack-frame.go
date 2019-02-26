package main

import "fmt"

func double(x int) int {
	return x * 2
}

func doubleV2(x int) {
	x *= 2
}

func main() {
	x := 20
	x = double(x)
	fmt.Println(x)

	x = 20
	doubleV2(x)
	fmt.Println(x)
}
