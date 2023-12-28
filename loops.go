package main

import "fmt"

func loops() {
	x := 0
	for x < 2 {
		fmt.Println("value of x :", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of x :", i)
	}

	names := []string{"john", "junior", "james"}
	for index, value := range names {
		fmt.Printf("%v is on index %v \n", value, index)
	}
}
