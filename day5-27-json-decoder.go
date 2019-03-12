package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	type Task struct {
		Desc string `json:"desc"`
		Done bool   `json:"done"`
	}
	r := strings.NewReader(`{"desc": "จ่ายค่าบัตรเครดิต","done":true}`)
	dec := json.NewDecoder(r)
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t)
}
