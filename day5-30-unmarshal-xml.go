package main

import (
	"encoding/xml"
	"fmt"
	"log"
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
	b := []byte(`<tasks><task done="true">จ่ายค่าบัตรเครดิต</task><task done="false">จ่ายค่าไฟ</task></tasks>`)
	var root Root
	err := xml.Unmarshal(b, &root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", root.Tasks)
}
