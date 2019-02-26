package main

import "fmt"

func main() {
	var v interface{}
	v = "World"
	switch d := v.(type) {
	case string:
		fmt.Println("String:", d)
		fmt.Println("Hello " + d)
	case int:
		fmt.Println("Int:", d)
		fmt.Println("Double of", d, "is", d*2)
	}
}
