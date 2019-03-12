package main

import (
	"encoding/json"
	"log"
	"os"
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
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(t)
	if err != nil {
		log.Fatal(err)
	}
}
