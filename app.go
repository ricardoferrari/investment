package main

import (
	"fmt"
	"math"
)

func main() {
	pv := getUserInput("Enter the present value: ")
	r := getUserInput("Enter the rate: ")
	n := getUserInput("Enter the period: ")

	fv, interest := calculateFutureValue(pv, r, n)

	printResults(fv, interest)
}

func printResults(fv, interest float64) {
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
