package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	count() // go count()
}
