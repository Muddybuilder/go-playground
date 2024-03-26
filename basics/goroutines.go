package main

import (
	"log"
	"os"
	"path/filepath"
)

func SearchFiles(dir string, lookFor string, ch chan string) {
	log.Println("[SEARCHING] " + dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(dir+file.Name(), file.IsDir())
		if file.Name() == lookFor {
			ch <- "[FOUND] " + filepath.Join(dir, file.Name())
			return
		}

	}
	ch <- "[NOT FOUND] " + dir
}

func main() {
	ch := make(chan string)
	go SearchFiles("./", "errors.go", ch)
	go SearchFiles("./", "errors1.go", ch)
	go SearchFiles("./", "errors32.go", ch)
	go SearchFiles("./", "errors44.go", ch)
	res := ""
	for i := 0; i < 4; i++ {
		res = <-ch
		log.Println(res)
	}
}
