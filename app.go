package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare a variable of type float64
	var pv float64
	var r float64 = 0.1
	var n float64 = 1

	// Prompt the user
	fmt.Print("Enter the present value: ")
	fmt.Scan(&pv)
	// Prompt the user
	fmt.Print("Enter the rate: ")
	fmt.Scan(&r)
	// Prompt the user
	fmt.Print("Enter the period: ")
	fmt.Scan(&n)

	r = r / 100
	// Calculate the future value
	fv := pv * math.Pow((1+r), n)

	// Display the future value
	fmt.Println("The future value is: ", fv)
}
