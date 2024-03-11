package main

import "fmt"

func main() {
	answer := ""

	for true {
		fmt.Print("Type command: ")
		fmt.Scan(&answer)
		if answer == "quit" {
			break
		} else if answer == "print" {
			fmt.Println("printing file")
		}
	}

	fmt.Println("program exit")

}
