package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main()  {
	//// #1 just place value with order, there will be a problem if you suddenly change the order
	////alex := person{"alex", "Anderson"}
	//
	//// #2 this is better, clearly and stable
	////alex := person{firstName: "alex", lastName: "Anderson"}
	//
	//// #3 init the empty struct
	//var alex = person{}
	//
	//// declare
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	//// %+v will print all different alex props
	//fmt.Printf("%+v", alex)

	gin := person{
		firstName: "Gin",
		lastName:  "Tonic",
		contact: contactInfo{
			email:   "gin_tonic@test.com",
			zipCode: 23845,
		},
	}

	fmt.Println(*&gin)

	//ginPointer := &gin
	// still woring without declare gin is pointer
	// because the pointerToPerson is pointer type, it will convert to pointer
	gin.updateName("jimmyexx")

	fmt.Printf("%+v", gin)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}