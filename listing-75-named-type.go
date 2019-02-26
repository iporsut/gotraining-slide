package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type A []interface{}
	type O map[string]interface{}
	posts := []O{
		{
			"id":    1,
			"title": "How to write Go?",
			"tags":  A{"Go", "Golang", "Tutorial"},
		},
		{
			"id":    2,
			"title": "What's pointer",
			"tags":  A{"Go", "Golang", "Pointer"},
		},
	}
	b, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
