package main

import "fmt"

func main() {
	ss := [3]string{
		"ประหยัด",
		"ชัดช่า",
		"อภิสิบ",
	}

	for i, v := range ss {
		fmt.Println(i, v)
	}
}
