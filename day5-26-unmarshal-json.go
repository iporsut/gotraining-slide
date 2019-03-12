package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Task struct {
		Desc string `json:"desc"`
		Done bool   `json:"done"`
	}
	b := []byte(`{"desc": "จ่ายค่าบัตรเครดิต","done":true}`)
	var t Task
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t)
}
