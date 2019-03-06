package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PARSE BEGIN OMIT
var listTmpl = template.Must(template.ParseFiles(filepath.Join("tmpl", "list-post.tmpl.html")))

// PARSE END OMIT

func main() {
	type Post struct {
		ID    int
		Title string
	}

	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		posts := []Post{
			{
				ID:    1,
				Title: "How to work with Go",
			},
			{
				ID:    2,
				Title: "How to use Go template",
			},
		}
		err := listTmpl.Execute(w, posts)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8000", nil)
}
