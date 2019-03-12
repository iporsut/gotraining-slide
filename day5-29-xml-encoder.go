package main

import (
	"encoding/xml"
	"log"
	"os"
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
	root := Root{
		Tasks: []Task{
			{
				Desc: "จ่ายค่าบัตรเครดิต",
				Done: true,
			},
			{
				Desc: "จ่ายค่าไฟ",
				Done: false,
			},
		},
	}
	enc := xml.NewEncoder(os.Stdout)
	err := enc.Encode(&root)
	if err != nil {
		log.Fatal(err)
	}
}
