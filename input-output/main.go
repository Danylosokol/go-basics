package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)

var url = "https://frontendmaster.com"

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func main() {
	defer fmt.Println("Buy!")
	var message string = "Hello from the module"
	var price float32 = 34.3
	newPrice := 52.2
	const pi byte = 3
	fmt.Println(message, price, newPrice, pi, url, data.MaxSpeed)
	fmt.Println("New line")
	stateTax, cityTax := calculateTax(100)
	stateTax, _ = calculateTaxWithNames(stateTax + cityTax)
	fmt.Println(stateTax, cityTax)
	printData()

	age := 2
	birthday(&age)
	fmt.Println(age)
}

func birthday(age *int) {
	if *age > 120 {
		panic("Too old to be true")
	}
	fmt.Printf("The pointer is %v and the value is %v\n", age, *age)
	*age++
}

func calculateTaxWithNames(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}
