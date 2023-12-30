package main

import (
	"fmt"
	"strconv"
)

func createBill() bill {
	name, _ := getInput("New Bill name : ")

	bl := newBill(name)
	return bl
}

func promptOptions(bl bill) {
	option, _ := getInput("choose option: \n a - add item \n s - save bill \n t - add tip \nEnter option: ")

	switch option {
	case "a":
		name, _ := getInput("Item Name :")
		price, _ := getInput("Item Price :")
		itemPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price value")
			promptOptions(bl)
		}
		bl.addItem(name, itemPrice)
		promptOptions(bl)
	case "s":
		bl.save()
	case "t":
		tip, _ := getInput("Add tip :")
		billTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Invalid tip value")
			promptOptions(bl)
		}
		bl.updateTip(billTip)
		promptOptions(bl)
	default:
		fmt.Println("invalid option")
		promptOptions(bl)
	}
}

func main() {
	bl := createBill()
	promptOptions(bl)
	fmt.Println(bl.format())
}
