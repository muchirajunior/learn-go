package main

import (
	"fmt"
	"math"
)

func sayGreenting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func cycleNames(names []string, function func(string)) {
	for _, value := range names {
		function(value)
	}
}

func areaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	cycleNames([]string{"john", "james", "mary"}, sayGreenting)

	number := areaOfCircle(10)

	fmt.Printf("number is %v \n", number)

	fmt.Println(number > 20)

	if number < 20 {
		fmt.Println("number is less than 20")
	} else {
		fmt.Println("number is greater than 20")
	}
}
