package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers []int

	// use for loop to set slice value
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(strconv.Itoa(n) + " is even")
		} else
		{
			fmt.Println(strconv.Itoa(n) + " is odd")
		}
	}
}
