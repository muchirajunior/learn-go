package main

import "fmt"

func update(item *string) {
	*item = "muchira"
}

func main() {
	bill1 := newBill("junior bill")
	fmt.Println(bill1)

}
