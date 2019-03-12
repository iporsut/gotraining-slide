package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	r := bytes.NewReader([]byte("some io.Reader stream to be read\n"))
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
