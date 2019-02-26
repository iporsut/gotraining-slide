package main

import "fmt"

// BEGIN TYPE DEFINE OMIT
type satang = int

type Customer struct {
	ID      int
	Name    string
	Balance satang
}

// END TYPE DEFINE OMIT

func main() {
	var c Customer
	fmt.Println(c)

	c = Customer{
		ID:      1,
		Name:    "Somrak",
		Balance: 100000,
	}
	fmt.Println(c)
	fmt.Println(c.ID, c.Name, c.Balance)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	c.Balance += 500000
	fmt.Printf("%#v\n", c)
}
