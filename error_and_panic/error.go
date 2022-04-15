package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	_, err := os.Open("Unknown.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The file does not exists!!")
	}
}