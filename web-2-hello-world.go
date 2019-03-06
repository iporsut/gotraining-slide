package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		fmt.Fprintf(w, `<!doctype html><html><body><h1>Hello %s</h1></body></html>`, name)
	})
	http.ListenAndServe(":8000", nil)
}
