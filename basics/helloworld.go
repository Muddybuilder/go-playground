package main

import "fmt"

func main() {
	// 1. hello world!
	fmt.Println("Hello, World!")

	/*2. Variables*/
	// explicit
	var firstName string
	firstName = "John"
	fmt.Println(firstName)

	// Implicit
	lastName := "Smith"
	fmt.Println(lastName)

	// Group definition
	var (
		balance  = 12
		custName = "Trump"
	)
	// string interpolation
	fmt.Printf("Customer %s has $%d on their account.", custName, balance)

}
