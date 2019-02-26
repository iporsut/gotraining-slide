package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("It's not a number")
		} else {
			// && เป็น operator and , || เป็น operator or
			if num%3 == 0 && num%5 == 0 {
				fmt.Println("FizzBuzz")
			} else if num%3 == 0 {
				fmt.Println("Fizz")
			} else if num%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(num)
			}
		}
	}
}
