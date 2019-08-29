package main

import "fmt"

func main() {

	// #1
	//var colors map[string]string

	// #2
	colors := make(map[string]string)

	// #3
	//colors := map[string]string{
	//	"red" : "#ff0000",
	//	"green" : "#4bf745",
	//}

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}
