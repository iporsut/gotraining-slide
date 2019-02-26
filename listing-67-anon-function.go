package main

import "fmt"

func do(op func(float64, float64) float64, x float64, y float64) {
	fmt.Println(op(x, y))
}

func main() {
	op := func(x float64, y float64) float64 {
		return x + y
	}
	fmt.Println(op(12.5, 13.7))

	do(func(x float64, y float64) float64 {
		return x * y
	}, 12.5, 13.7)
}
