package main

import (
	"fmt"
	"strconv"
)

func main() {
	log("Hi")

	sum, mult := add_and_mult(1, 3)

	log(strconv.Itoa(sum))
	log(strconv.Itoa(mult))

}

func log(message string) {
	fmt.Println(message)
}

func add_and_mult(num1 int, num2 int) (sum int, product int) {
	sum = num1 + num2
	product = num1 * num2
	return
}
