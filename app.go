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
	pv = getUserInput("Enter the present value: ")
	// Prompt the user
	r = getUserInput("Enter the rate: ")
	// Prompt the user
	n = getUserInput("Enter the period: ")

	// Calculate the future value
	fv, interest := calculateFutureValue(pv, r, n)

	formattedFVString := fmt.Sprintf("The future value is %.2f", fv)
	fmt.Println(formattedFVString)

	fmt.Printf("The interest is %.2f\n", interest)

}

func getUserInput(info string) float64 {
	var value float64
	fmt.Print(info)
	fmt.Scan(&value)
	return value
}

func calculateFutureValue(pv, r, n float64) (fv float64, interest float64) {
	r = r / 100
	// Calculate the future value
	fv = pv * math.Pow((1+r), n)
	// Deduct the present value from the future value
	interest = fv - pv
	// return fv, interest
	return
}
