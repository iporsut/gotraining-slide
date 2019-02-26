package main

import "fmt"

func main() {
	ss := [3]string{
		"ประหยัด",
		"ชัดช่า",
		"อภิสิบ",
	}

	for i := range ss {
		fmt.Println(ss[i])
	}
}
