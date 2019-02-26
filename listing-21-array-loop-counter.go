package main

import "fmt"

func main() {
	ss := [3]string{
		"ประหยัด",
		"ชัดช่า",
		"อภิสิบ",
	}

	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}
}
