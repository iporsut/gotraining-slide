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

// BEGIN NEW CUSTOMER OMIT
func NewCustomer(name string, balance satang) *Customer {
	return &Customer{
		ID:      1, // TODO: Auto increment ID
		Name:    name,
		Balance: balance,
	} // Move to heap
}

func main() {
	c := NewCustomer("Somrak", 100000)
	fmt.Println(c)

	fmt.Println(c)
	fmt.Println(c.ID, c.Name, c.Balance)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	c.Balance += 500000
	fmt.Printf("%#v\n", c)
}

// END NEW CUSTOMER OMIT
