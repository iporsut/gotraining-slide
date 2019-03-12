package main

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, World")
	if _, err := io.Copy(ioutil.Discard, r); err != nil {
		log.Fatal(err)
	}
}
