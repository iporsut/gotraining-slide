package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	customers := []map[string]interface{}{
		map[string]interface{}{
			"id":      1,
			"name":    "Somrak",
			"balance": 1000,
		},
		map[string]interface{}{
			"id":      2,
			"name":    "Somsri",
			"balance": 1500.0,
		},
	}

	b, err := json.Marshal(customers) // b, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b)) // or fmt.Printf("%s\n", b)
}
