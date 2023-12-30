package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	total float64
	tip   float64
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

	fs += fmt.Sprintf("%-25v ... $%f\n", "tip:", bl.tip)

	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", bl.total)

	return fs
}

// update tip
func (bl *bill) updateTip(tip float64) {
	bl.tip = tip
}

func (bl *bill) addItem(name string, price float64) {
	bl.items[name] = price
}

func (bl *bill) save() {
	data := []byte(bl.format())
	err := os.WriteFile("data/"+bl.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved bill succesfully")
}
