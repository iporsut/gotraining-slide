package main

import (
	"fmt"
	"log"
	"os"
)

type BinaryOperatorError struct {
	LeftOperand  float64
	RightOperand float64
	Operator     string
	Msg          string
}

func (err BinaryOperatorError) Error() string {
	return fmt.Sprintf("error: %s when %f %s %f", err.Msg, err.LeftOperand, err.Operator, err.RightOperand)
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, BinaryOperatorError{a, b, "/", "Cannot divide by zero"}
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
