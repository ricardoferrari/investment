package main

import (
	"testing"
)

func TestCalculateFutureValue(t *testing.T) {
	tests := []struct {
		pv       float64
		r        float64
		n        float64
		expected float64
		interest float64
	}{
		{1000, 10, 1, 1100, 100},
		{1000, 5, 2, 1102.5, 102.5},
		{1000, 1, 12, 1126.8250301319697, 126.8250301319697},
	}

	for _, test := range tests {
		fv, interest := calculateFutureValue(test.pv, test.r, test.n)
		if fv != test.expected || interest != test.expected-test.pv {
			t.Errorf("For pv=%.2f, r=%.2f, n=%.2f; expected %.2f but got %.2f", test.pv, test.r, test.n, test.expected, fv)
		}
	}
}

func TestGetUserInput(t *testing.T) {
	// This function is interactive and requires user input, which is not suitable for automated testing.
	// You can mock this function if needed.
}

func TestPrintResults(t *testing.T) {
	// This function is just a wrapper for fmt.Println and fmt.Printf, which are already tested by the standard library.
	// No need to test this function.
}
