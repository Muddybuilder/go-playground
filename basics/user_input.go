package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	players := []string{}
	fmt.Print("Enter number of players: ")
	fmt.Scan(&input)

	playerNumber, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	for i := 1; i <= playerNumber; i++ {
		fmt.Printf("Enter Player %d name: ", i)
		var newPlayer string
		fmt.Scan(&newPlayer)
		players = append(players, newPlayer)
	}

	fmt.Print("Players are: ")
	for _, player := range players {
		fmt.Print(player + " ")
	}
}
