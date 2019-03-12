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
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
