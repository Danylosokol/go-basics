package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Firtman"} // new Instructor
	max.FirstName = "Maximiliano"
	print(max.Print())

	kyle := data.NewInstructor("Kyle", "Simpson")
	print(kyle.Print())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}
	fmt.Printf("%v", goCourse)

	swiftWS := data.NewWorkshop("Swift with iOS", max)
	fmt.Printf("%v", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
