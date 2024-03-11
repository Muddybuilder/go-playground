package main

import "fmt"

func main() {
	count := 10
	// for statement
	for i := 0; i < count; i++ {

	}

	// while statement
	keepGoing := true
	for keepGoing {
		magicWord := "Exit"
		fmt.Printf("Type %s ro exit: ", magicWord)
		var ans string
		fmt.Scan(&ans)

		if ans == magicWord {
			keepGoing = false
		}

	}

	// for-each loop
	v := []int{1, 2, 3}
	for i, v := range v {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
