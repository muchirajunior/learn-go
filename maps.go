package main

import "fmt"

func maps() {
	person := map[string]float64{
		"age":    22,
		"number": 9,
		"height": 90,
	}

	fmt.Println(person["age"])

	for key, value := range person {
		fmt.Println(key, " - ", value)
	}
}
