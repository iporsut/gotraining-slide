package main

import "fmt"

func concatKV(m map[string]string, k1, k2 string) {
	m[k1+"-"+k2] = m[k1] + m[k2]
}

func main() {
	m := map[string]string{
		"firstName": "Somsak",
		"lastName":  "Sawasdee",
	}

	concatKV(m, "firstName", "lastName")

	fmt.Println(m)
}
