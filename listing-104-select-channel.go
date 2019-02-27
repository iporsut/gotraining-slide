package main

import (
	"context"
	"fmt"
	"time"
)

func sender(ctx context.Context, ch chan time.Time, d time.Duration) {
	for {
		select {
		case ch <- <-time.After(d):
		case <-ctx.Done():
			// stop goroutine
			// when have done signal from context
			return
		}
	}
}

func main() {
	chOneSec := make(chan time.Time)
	chFiveSec := make(chan time.Time)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	go sender(ctx, chOneSec, 1*time.Second)
	go sender(ctx, chFiveSec, 5*time.Second)
	for {
		select {
		case t := <-chOneSec:
			fmt.Println("From one second sender:", t)
		case t := <-chFiveSec:
			fmt.Println("From five second sender:", t)
		case <-ctx.Done():
			return
		}
	}
}
