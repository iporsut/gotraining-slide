package main

import "fmt"

func wrapOp(op func(float64, float64) float64) func(float64, float64) {
	return func(x float64, y float64) {
		fmt.Println(op(x, y))
	}
}

func main() {
	printOp := wrapOp(func(x float64, y float64) float64 {
		return x * y
	})
	printOp(20.5, 30.7)
}
