package main

import "fmt"

func main() {
	type Basket struct {
		title    string
		desc     string
		quantity int
		ppu      int
		total    int
	}

	var basketA Basket
	basketA.title = "LEGO set"
	basketA.desc = "4000 pieces"
	basketA.quantity = 1
	basketA.ppu = 600
	basketA.total = 660

	basketB := Basket{
		"Plushy",
		"Plushy toy",
		3,
		5,
		15,
	}

	list := make([]Basket, 0)

	list = append(list, basketA)
	list = append(list, basketB)
	total := 0

	for i := 0; i < len(list); i++ {
		total += list[i].total
	}

	fmt.Println("Total price is: ", total, "GBP")

}
