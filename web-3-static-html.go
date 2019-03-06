package main

import "net/http"

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
