package main

import "fmt"

func main() {
	ss := [3]string{
		"ประหยัด",
		"ชัดช่า",
		"อภิสิบ",
	}

	for _, v := range ss {
		fmt.Println(v)
	}
}
