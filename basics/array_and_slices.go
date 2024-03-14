package main

import "fmt"

func main() {
	var response string
	respArray := make([]string, 0)

	for {
		fmt.Print("command> ")
		fmt.Scan(&response)
		if response == "quit" {
			break
		} else if response == "new" {
			fmt.Print("Entry:")
			fmt.Scan(&response)
			respArray = append(respArray, response)

			fmt.Println("saving entry")
		} else if response == "list" {
			fmt.Println("Listing entries")
			for _, v := range respArray {
				fmt.Println(v)
			}
		} else {
			fmt.Println("Unknown command", response)
		}
	}
	fmt.Println("Byeee")
}
