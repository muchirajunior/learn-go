package main

import (
	"fmt"
	"strings"
)

func variables() {
	var nameOne = "muchira"
	var number int8 = 10
	var decimal float32 = 283
	var bigDecimal float64 = 66.09373
	var boolean bool = false
	nameTwo := "junior"
	fmt.Println(nameOne, nameTwo, number, decimal, bigDecimal, boolean)

	fmt.Println("my full name is", nameOne, nameTwo)

	//formated string
	fmt.Printf("FullName : %v - %v \n", nameOne, nameTwo)
	fmt.Printf("FullName : %q - %q\n", nameOne, nameTwo)
	fmt.Printf("number is type %T\n", number)
	fmt.Printf("floting number %0.2f \n", bigDecimal)

	//save formated string
	formatedString := fmt.Sprintf("number is %v", number)
	fmt.Println(formatedString)

	//arrays
	var numbers [4]int = [4]int{1, 2, 3}
	numbers[3] = 100
	fmt.Println(numbers, len(numbers))

	//slices
	var scores = []int{20, 40, 80}
	scores[1] = 90
	scores = append(scores, 88)
	fmt.Println(scores, len(scores))

	//strings
	hello := "hellooooooo"
	fmt.Println(strings.Contains(hello, "ll"))
	fmt.Println(strings.ReplaceAll(hello, "o", "x"))
}
