package main

import "fmt"

func main() {
	var wc map[string]int // wc is nil
	fmt.Println(wc["nonexists"])

	wc = make(map[string]int)
	fmt.Println(wc["nonexists"])

	wc["hello"]++
	c := wc["hello"]
	fmt.Println(c)

	c, ok := wc["nonexists"]
	fmt.Println(c, ok)

	c, ok = wc["hello"]
	fmt.Println(c, ok)
}
