package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	total float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"cake":  800,
			"juice": 20,
		},
		total: 0,
	}

	return b
}

// format struct
func (bl bill) format() string {
	fs := "Bill info : \n"
	for key, value := range bl.items {
		fs += fmt.Sprintf("%-25v ..$ %v \n", key+":", value)
		bl.total += value
	}

	fs += fmt.Sprintf("%-25v ....$%0.2f", "totol:", bl.total)

	return fs
}
