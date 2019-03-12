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

	t := &Task{
		Desc: "จ่ายค่าบัตรเครดิต",
		Done: true,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
