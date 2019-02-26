package main

import (
	"fmt"
	"time"
)

func main() {
	const timeFormatLayout = "15:04:05 02/01/2006"
	for i := 1; i <= 10; i++ {
		now := time.Now()
		fmt.Printf("%d : %s\n", i, now.Format(timeFormatLayout))
		time.Sleep(1 * time.Second)
	}
}
