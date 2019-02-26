package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseToken(expr string) (float64, string, float64) {
	tokens := strings.Fields(strings.TrimSpace(expr))
	if len(tokens) != 3 {
		panic("expression is not a binary")
	}
	left, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		panic("left operand is not a number")
	}
	op := tokens[1]
	right, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		panic("right operand is not a number")
	}
	return left, op, right
}

func binaryEval(expr string) float64 {
	switch left, op, right := parseToken(expr); op {
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
	fmt.Println(binaryEval("99 + 11"))
}
