package main

import (
	"fmt"
	"sort"
)

type satang = int

type Customer struct {
	ID      int
	Name    string
	Balance satang
}

func main() {
	// BEGIN DATA OMIT
	custs := []Customer{
		Customer{
			ID:      1,
			Name:    "Somrak",
			Balance: 100000,
		},
		Customer{
			ID:      2,
			Name:    "Somsri",
			Balance: 500000,
		},
		Customer{
			ID:      3,
			Name:    "Anak",
			Balance: 1000000,
		},
	}
	// END DATA OMIT
	// BEGIN SORT OMIT
	sort.Slice(custs, func(i, j int) bool {
		return custs[i].Balance > custs[j].Balance
	})
	fmt.Printf("%#v\n", custs)
	// END SORT OMIT
}
