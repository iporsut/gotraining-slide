package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PARSE BEGIN OMIT
var helloTmpl = template.Must(template.ParseFiles(filepath.Join("tmpl", "hello.tmpl.html")))

// PARSE END OMIT

func main() {
	type HelloPage struct {
		Title string
		Name  string
	}

	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		err := helloTmpl.Execute(w, HelloPage{
			Title: "Hello",
			Name:  r.FormValue("name"),
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8000", nil)
}
