package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f2, err := os.Create("data-2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(f2, f); err != nil {
		log.Fatal(err)
	}
}
