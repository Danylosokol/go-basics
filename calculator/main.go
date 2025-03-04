package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")

	fmt.Println("Which operation do you want to perform? (add, substract, multiply, divide)")
	fmt.Scanf("%s", &operation)

	fmt.Println("Enter the first number:")
	fmt.Scanf("%v", &number1)

	fmt.Println("Enter the second number:")
	fmt.Scanf("%v", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}
