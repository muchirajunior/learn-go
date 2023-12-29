package main

import (
	"fmt"
)

func createBill() bill {
	name, _ := getInput("New Bill name : ")

	bl := newBill(name)
	return bl
}

func promptOptions(bl bill) {
	option, _ := getInput("choose option: \n a - add item \n s - save bill \n t - add tip \nEnter option: ")
	fmt.Println(option)
}

func main() {
	bl := createBill()
	promptOptions(bl)
	fmt.Println(bl.format())
}
