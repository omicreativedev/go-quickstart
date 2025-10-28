package main

import (
	"fmt"
	"strings"
)

// multiplies two numbers and returns the output
func multiply(a, b int) int {
	return a * b
}

// recursive function that calculates a factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// splits a string into two strings and returns both
func splitString(s string) (string, string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// tests whether Go is pass-by-reference or pass-by-value
func passByTest(s string, n int) (string, int) {
	s = "modified"
	n = 999
	return s, n
}

func main() {
	// Call functions and save results
	product := multiply(5, 7)
	fact := factorial(5)
	left, right := splitString("hello world")
	
	originalStr := "original"
	originalNum := 123
	modifiedStr, modifiedNum := passByTest(originalStr, originalNum)
	
	// Print results
	fmt.Printf("Multiply: %d\n", product)
	fmt.Printf("Factorial: %d\n", fact)
	fmt.Printf("Split: %q, %q\n", left, right)
	fmt.Printf("Pass-by test - Original: %q, %d | After: %q, %d\n", 
		originalStr, originalNum, modifiedStr, modifiedNum)
}