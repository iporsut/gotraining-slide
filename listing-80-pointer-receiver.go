package main

import (
	"fmt"
	"math"
)

type Satang = int

type Customer struct {
	ID        int
	Firstname string
	Lastname  string
	Balance   Satang
}

func (c Customer) Name() string {
	return c.Firstname + " " + c.Lastname // fmt.Sprint("%s %s", c.Firstname, c.Lastname)
}

func (c *Customer) EarnInterest(percent float64) {
	c.Balance = c.Balance + int(math.Round(float64(c.Balance)*percent/100))
}

func main() {
	c := Customer{
		ID:        1,
		Firstname: "Somrak",
		Lastname:  "Jaidee",
		Balance:   100000,
	}
	fmt.Println(c.ID, c.Name(), c.Balance)
	c.EarnInterest(2.5)
	fmt.Println(c.ID, c.Name(), c.Balance)
}
