package main

import "fmt"

type Satang int

func (s Satang) ToBaht() float64 {
	return float64(s) / 100
}

func main() {
	c := Customer{
		ID:        1,
		Firstname: "Somrak",
		Lastname:  "Jaidee",
		Balance:   100000,
	}
	fmt.Println(c.ID, c.Name(), c.Balance.ToBaht())
}

type Customer struct {
	ID        int
	Firstname string
	Lastname  string
	Balance   Satang
}

func (c Customer) Name() string {
	return c.Firstname + " " + c.Lastname // fmt.Sprint("%s %s", c.Firstname, c.Lastname)
}
