package main

import "fmt"

// BEGIN OMIT
func sender(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sender(ch)       // start sender goroutine
	for v := range ch { // receiver
		fmt.Println(v)
	}
	fmt.Println("DONE")
}

// END OMIT
