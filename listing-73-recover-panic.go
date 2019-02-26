package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered from panic", err)
		}
	}()
	panic("Error")
}
