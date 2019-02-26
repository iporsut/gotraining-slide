package main

import "fmt"

func main() {
	var m map[string]string // zero value is nil
	fmt.Println(m == nil)

	m = make(map[string]string) // no element map , m = map[string]string{}
	m["name"] = "Weerasak"
	m["job"] = "Senior Software Developer"

	fmt.Println(m)

	m = map[string]string{
		"name": "Weerask Chongnguluam",
		"job":  "Programmer",
	}

	fmt.Println(m["name"], m["job"])
}
