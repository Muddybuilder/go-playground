package main

import (
	"fmt"
	"strings"
)

func main() {
	phonebook := make(map[string]string)

	fmt.Println("Welcome to your phonebook.")
	var name string
	var number string
	var cmd string
	for true {
		fmt.Print("Command> ")
		fmt.Scan(&cmd)
		if cmd == "store" {
			fmt.Print("Enter contact: ")

			fmt.Scan(&name, &number)

			phonebook[name] = number
			fmt.Println("Contact saved")
		} else if cmd == "list" {
			if len(phonebook) == 0 {
				fmt.Println("Empty phonebook!")
			} else {
				for k, v := range phonebook {
					fmt.Println(k, v)
				}
			}
		} else if cmd == "lookup" {
			fmt.Print("Enter name: ")
			fmt.Scan(&cmd)
			val, exist := phonebook[cmd]
			if exist {

				fmt.Println(cmd, "has number: ", val)

			} else {
				var yesNo string
				fmt.Println("Contact doesn't exist, do you want to add it? y/n:")
				fmt.Scan(&yesNo)
				if strings.ToLower(yesNo) == "y" || strings.ToLower(yesNo) == "yes" {
					fmt.Print("Enter contact: ")

					fmt.Scan(&name, &number)

					phonebook[name] = number
					fmt.Println("Contact saved")

				}

			}

		} else if cmd == "quit" {
			fmt.Println("BYEEE")
			break
		} else {
			fmt.Println("Unknown command")

		}

	}

}
