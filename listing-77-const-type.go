package main

import "fmt"

type Size int

const (
	SSS Size = iota
	SS
	S
	M
	L
	XL
	XXL
)

var sizeStr = map[Size]string{
	SSS: "SSS",
	SS:  "SS",
	S:   "S",
	M:   "M",
	L:   "L",
	XL:  "XL",
	XXL: "XXL",
}

// func (s Size) String() string {
// 	return sizeStr[s]
// }

func main() {
	fmt.Printf("%T, %v\n", XXL, XXL)
}
