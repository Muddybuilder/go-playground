package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Id    int    `json: "id"`
	Items []Item `json: "items`
}

type Item struct {
	Id       int     `json: "id"`
	Quantity int     `json: "quantity"`
	Total    float32 `json: "total"`
}

func main() {
	file, _ := os.ReadFile("response.json")
	data := Orders{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Orders); i++ {
		fmt.Printf("Order ID: %d", data.Orders[i].Id)

		for j := 0; j < len(data.Orders[i].Items); j++ {
			fmt.Printf("Item Id: %d\n", data.Orders[i].Items[j].Id)
			fmt.Printf("Item Quantity: %d\n", data.Orders[i].Items[j].Quantity)
			fmt.Printf("Item Total: %.2f\n", data.Orders[i].Items[j].Total)

		}
	}

}
