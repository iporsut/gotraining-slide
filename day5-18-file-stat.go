package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Stat("data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Is dir?: %t\n", f.IsDir())
	fmt.Printf("Size: %d\n", f.Size())
	fmt.Printf("Mode: %s\n", f.Mode())
}
