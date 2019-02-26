package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
}
