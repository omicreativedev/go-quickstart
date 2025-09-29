package main

import (
	"fmt"
	"reflect"
)

func main() {
	//
	// Declarations: Implicit
	//
	
	// Integers
	x1 := 15
	x2 := 8
	
	// Floating-point numbers
	y1 := 3.5
	y2 := 4.5
	
	// Boolean values
	isTrue, isFalse := true, false
	
	// Strings
	myName := "Salome"
	myType := "female"
	
	// Arrays
	country := [...]string{"Spanish", "German", "French", "English", "Canadian"}
	id := [5]int{3, 7, 5, 4, 8}

	// Slice (dynamic array)
	prices := []float64{12.5, 7.3, 9.99, 15.0}
	
	// Map (dictionary)
	friends := map[string]string{
		"Joe": "male",
		"Kay": "female",
		"Sam": "nonbinary",
	}

	// Print variables and their types
	fmt.Println("\nVariables and Types\n")
	fmt.Println("x1 is", x1, "of the type", reflect.TypeOf(x1))
	fmt.Println("x2 is", x2, "of the type", reflect.TypeOf(x2))
	fmt.Println("y1 is", y1, "of the type", reflect.TypeOf(y1))
	fmt.Println("y2 is", y2, "of the type", reflect.TypeOf(y2))
	fmt.Println("isTrue is", isTrue, "of the type", reflect.TypeOf(isTrue))
	fmt.Println("isFalse is", isFalse, "of the type", reflect.TypeOf(isFalse))
	fmt.Println("myName is", myName, "of the type", reflect.TypeOf(myName))
	fmt.Println("myType is", myType, "of the type", reflect.TypeOf(myType))
	fmt.Println("country is", country, "of the type", reflect.TypeOf(country))
	fmt.Println("id is", id, "of the type", reflect.TypeOf(id))
	fmt.Println("prices is", prices, "of the type", reflect.TypeOf(prices))
	fmt.Println("friends is", friends, "of the type", reflect.TypeOf(friends))
	
	//
	// Arithmetic Operations (Same Types)
	//
	fmt.Println("\nArithmetic Operations (Same Types)\n")
	fmt.Println("x1 + x2 =", x1+x2)
	fmt.Println("y2 - y1 =", y2-y1)
	fmt.Println("x1 * x2 =", x1*x2)
	fmt.Println("x1 / x2 =", x1/x2) // integer division
	fmt.Println("y2 / y1 =", y2/y1) // float division
	
	//
	// Arithmetic Operations (Different Types)
	//
	fmt.Println("\nArithmetic Operations (Different Types)\n")
	// convert int to float64 for addition
	fmt.Println("x1 + y1 =", float64(x1)+y1) 
	fmt.Println("x2 * y2 =", float64(x2)*y2)
	
	//
	// Type Casting Examples
	//
	fmt.Println("\nType Casting Examples\n")
	fmt.Println("Convert y1 to int:", int(y1))
	fmt.Println("Convert x1 to float64:", float64(x1))
	
	//
	// String Concatenation
	//
	fmt.Println("\nString Operations\n")
	greeting := "Hello, " + myName
	fmt.Println(greeting)
	
	//
	// Array and Slice Operations
	//
	fmt.Println("\nArray and Slice Operations\n")
	fmt.Println("First country:", country[0])
	fmt.Println("All prices:", prices)
	prices = append(prices, 20.99)
	fmt.Println("Prices after append:", prices)
	
	//
	// Map Operations
	//
	fmt.Println("\nMap Operations\n")
	fmt.Println("Sam's gender is", friends["Sam"])
	friends["Alex"] = "nonbinary"
	fmt.Println("Updated friends map:", friends)
	
	//
	// Limitations / Pitfalls in Go
	//
	fmt.Println("\nLimitations / Pitfalls\n")
	fmt.Println("1. Go requires explicit type conversions; mixing int and float without conversion fails.")
	fmt.Println("2. Arrays have fixed size, slices are preferred for dynamic collections.")
	fmt.Println("3. Implicit declarations only work with ':=' and must be used in the same scope.")
	fmt.Println("4. Maps cannot have duplicate keys; accessing a non-existent key returns zero value.")
}
