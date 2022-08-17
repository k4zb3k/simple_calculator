package main

import (
	"fmt"
)

var a, b, result float64
var oper string

	// Calculator that built with 'if' statement

	/*
	In order to check calculator that built with 'if' statement,
	please comment all code that consist in "func switchType()" and a "for {switchType()}" in main
	*/
func ifType() {
	fmt.Print("enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("enter an operator(+, -, *, /): ")
	fmt.Scanln(&oper)
	fmt.Print("enter second number: ")
	fmt.Scanln(&b)

	if oper == "+" {
		result = a + b
		fmt.Println("result is: ", result)
	} else if oper == "-" {
		result = a - b
		fmt.Println("result is: ", result)
	} else if oper == "*" {
		result = a * b
		fmt.Println("result is: ", result)
	} else if oper == "/" {
		result = a / b
		fmt.Println("result is: ", result)
	} else {
		fmt.Println("Error")
	}
}

	// Calculator that built with 'switch'

	/*
	In order to check calculator that built with 'if' statement,
	please comment all code that consist in "func ifType()" and a "for {ifType()}" in main
	*/
func switchType() {
	fmt.Print("enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("enter an operator(+, -, *, /): ")
	fmt.Scanln(&oper)
	fmt.Print("enter second number: ")
	fmt.Scanln(&b)

	switch oper {
	case "+":
		result = a + b
		fmt.Println("result: ", result)
	case "-":
		result = a - b
		fmt.Println("result: ", result)
	case "*":
		result = a * b
		fmt.Println("result: ", result)
	case "/":
		result = a / b
		fmt.Println("result: ", result)
		
	default:
		fmt.Println("Error")
	}
}

func main() {
	for {
		switchType()
		var command string
		fmt.Println("\nDo you want to continue? y/n")
		fmt.Scanln(&command)
		if command == "n" {
			break
		}
	   }

	for {
		ifType()
		var command string
		fmt.Print("\nDo you want to continue? y/n: ")
		fmt.Scanln(&command)
		if command == "n" {
			break
		}
	}
}