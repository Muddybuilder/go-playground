package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func main() {

	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)

	}

	log.SetOutput(f)

	log.Println("Starting Program...")
	no := Divide(10, 0)
	fmt.Println(no)

	no = Divide(10, 1)
	fmt.Println(no)
	f.Close()

}

func errorHandler() {
	r := recover()
	if r != nil {
		log.Println(r, string(debug.Stack()))

	}
}

func Divide(nom int, div int) float32 {
	defer errorHandler()
	if div == 0 {
		panic("can't divide by 0")
	}
	return float32(nom) / float32(div)
}
