package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

func main() {
	type Task struct {
		Desc string `xml:",chardata"`
		Done bool   `xml:"done,attr"`
	}

	type Root struct {
		XMLName xml.Name `xml:"tasks"`
		Tasks   []Task   `xml:"task"`
	}
	r := strings.NewReader(`<tasks><task done="true">จ่ายค่าบัตรเครดิต</task><task done="false">จ่ายค่าไฟ</task></tasks>`)
	dec := xml.NewDecoder(r)
	var root Root
	err := dec.Decode(&root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", root.Tasks)
}
