package main

import (
	"fmt"
	"time"
)

func main() {
	const timeFormatLayout = "15:04:05 02/01/2006"
	now := time.Now()
	for {
		fmt.Println(now.Format(timeFormatLayout))
		time.Sleep(1 * time.Second)
		now = time.Now()
	}
}
