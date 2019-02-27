package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// BEGIN OMIT
type WordCounter struct {
	m  map[string]int
	mu sync.Mutex
}

func NewWordCounter() *WordCounter {
	return &WordCounter{
		m: make(map[string]int),
	}
}

func (wc *WordCounter) Count(word string) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	wc.m[word]++
}

func (wc *WordCounter) Println() {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for w, c := range wc.m {
		fmt.Printf("%s: %d\n", w, c)
	}
}

// END OMIT

// COUNT OMIT
func count(wc *WordCounter, scanner *bufio.Scanner, wg *sync.WaitGroup) {
	defer wg.Done()
	for scanner.Scan() {
		for _, w := range strings.Fields(scanner.Text()) {
			wc.Count(w)
		}
	}
}

// END COUNT OMIT

func main() {
	wc := NewWordCounter()
	wg := &sync.WaitGroup{}
	files := []string{"file_one.txt", "file_two.txt"}
	for _, filename := range files {
		wg.Add(1)
		go func(filename string) {
			f, err := os.Open(filename)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer f.Close()
			count(wc, bufio.NewScanner(f), wg)
		}(filename)
	}
	wg.Wait()
	wc.Println()
}
