package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<!doctype html><html><body><h1>Hello</h1></body></html>`)
	})
	http.ListenAndServe(":8000", nil)
}
