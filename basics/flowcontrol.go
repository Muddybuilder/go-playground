package main

import (
	"fmt"
)

func main() {

	testScoreGrade5 := 80
	testScoreGrade4 := 60
	testScoreGrade3 := 50
	testScore := 49

	hasGas := true
	hasKeyIgnition := true

	hasBurger := true
	hasSandwich := false
	printMessage := true

	if printMessage {
		fmt.Println("Message")
	}

	if testScore >= testScoreGrade5 {
		fmt.Println("Top mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with distinction")

	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass with distinction")
	} else if testScore >= 0 {
		fmt.Println("Failed")
	} else {
		fmt.Println("Why did you get negative score???")
	}

	if hasGas && hasKeyIgnition {
		fmt.Println("Can drive!")
	}

	if hasBurger || hasSandwich {
		fmt.Println("Can eat!")
	}

	if !hasSandwich {
		fmt.Println("No sammich :(")
	}

}
