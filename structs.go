package main

type bill struct {
	name  string
	items map[string]float64
	total float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		total: 0,
	}

	return b
}
