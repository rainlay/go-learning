package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main()  {
	// #1 just place value with order, there will be a problem if you suddenly change the order
	//alex := person{"alex", "Anderson"}

	// #2 this is better, clearly and stable
	//alex := person{firstName: "alex", lastName: "Anderson"}

	// #3 init the empty struct
	var alex = person{}

	// declare
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	// %+v will print all different alex props
	fmt.Printf("%+v", alex)
}
