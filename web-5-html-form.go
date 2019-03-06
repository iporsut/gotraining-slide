package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/new/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		firstName := r.FormValue("firstname")
		lastName := r.FormValue("lastname")
		tmpl := "<html><body><h1>FirstName: %s, LastName: %s</h1></body></html>"
		fmt.Fprintf(w, tmpl, firstName, lastName)
	})

	http.ListenAndServe(":8000", nil)
}
