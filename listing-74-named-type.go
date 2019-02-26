package main

import "fmt"

type Distance float64

func area(w Distance, h Distance) Distance {
	return w * h
}

func main() {
	var w Distance = 10
	var h Distance = 20
	fmt.Println(area(w, h))

	i := 10
	j := 20
	fmt.Println(area(i, j)) // FIXED: fmt.Println(area(Distance(i), Distance(j)))
}
