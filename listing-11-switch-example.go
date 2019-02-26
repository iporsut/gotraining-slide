package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryEval(expr string) float64 {
	tokens := strings.Fields(strings.TrimSpace(expr))
	if len(tokens) != 3 {
		panic("expression is not a binary")
	}

	left, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		panic("left operand is not a number")
	}

	right, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		panic("right operand is not a number")
	}

	switch tokens[1] {
	default:
		panic("not supported operator")
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
}

func main() {
	fmt.Println(binaryEval("1.5 + 2.5"))
}
