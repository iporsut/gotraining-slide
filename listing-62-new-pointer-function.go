package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(p)
	*p = 20
	fmt.Println(*p)
}
