package main

import "fmt"

func main() {
	bill1 := newBill("junior bill")
	bill1.updateTip(20)
	fmt.Println(bill1.format())

}
