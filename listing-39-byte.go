package main

import "fmt"

func main() {
	var b []byte = []byte("ก")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d", b[i])
	}
	fmt.Println()
}
