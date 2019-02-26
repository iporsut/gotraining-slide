package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(x, y float64) float64 {
	return x + y
}
func sub(x, y float64) float64 {
	return x - y
}
func mul(x, y float64) float64 {
	return x * y
}
func div(x, y float64) float64 {
	return x / y
}

func parseToken(expr string) (left float64, op string, right float64) {
	var err error
	tokens := strings.Fields(strings.TrimSpace(expr))
	if len(tokens) != 3 {
		panic("expression is not a binary")
	}
	left, err = strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		panic("left operand is not a number")
	}
	op = tokens[1]
	right, err = strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		panic("right operand is not a number")
	}
	return
}

// BEGIN SELECT OP OMIT
func selectOp(op string) func(float64, float64) float64 {
	switch op {
	default:
		panic("not supported operator")
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	}
}

func binaryEval(expr string) float64 {
	left, op, right := parseToken(expr)
	fn := selectOp(op)
	return fn(left, right)
}

// END SELECT OP OMIT
func main() {
	fmt.Println(binaryEval("99 + 11"))
}
