package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	total float64
	tip   int
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		total: 0,
	}

	return b
}

// format struct
func (bl bill) format() string {
	fs := bl.name + " : \n"
	for key, value := range bl.items {
		fs += fmt.Sprintf("%-25v .. $ %v \n", key+":", value)
		bl.total += value
	}

	fs += fmt.Sprintf("%-25v ... %v\n", "tip:", bl.tip)

	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", bl.total)

	return fs
}

// update tip
func (bl *bill) updateTip(tip int) {
	bl.tip = tip
}

func (bl *bill) addItem(name string, price float64) {
	bl.items[name] = price
}
