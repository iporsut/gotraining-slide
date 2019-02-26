package main

import "fmt"

func main() {
	var (
		order1       = 20
		order2       = 30
		order3       = 50
		discountRate = 10 // 10 percent
		people       = 10
	)

	total := float64(order1 + order2*2 + order3*3)
	discount := total * float64(discountRate) / 100.0
	average := (total - discount) / float64(people)

	fmt.Println(average)
}
