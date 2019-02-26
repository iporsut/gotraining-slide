package main

import "fmt"

type satang = int
type Customer struct {
	ID        int
	Firstname string
	Lastname  string
	Balance   satang
}

func (c Customer) Name() string {
	return c.Firstname + " " + c.Lastname // fmt.Sprint("%s %s", c.Firstname, c.Lastname)
}
func main() {
	c := Customer{
		ID:        1,
		Firstname: "Somrak",
		Lastname:  "Jaidee",
		Balance:   100000,
	}
	fmt.Println(c.ID, c.Name(), c.Balance)
}
