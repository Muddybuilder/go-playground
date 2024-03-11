package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		panic("Please give two arguments.")
	}

	num1, err1 := strconv.Atoi(os.Args[1])

	if err1 != nil {
		panic("Invalid input!!")
	}

	num2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic("Invalid input!!")
	}

	fmt.Printf("Sum: %d", num1+num2)

}
