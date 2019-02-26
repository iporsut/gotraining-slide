package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	r, err := Div(10, 0)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	fmt.Println(r)
}
