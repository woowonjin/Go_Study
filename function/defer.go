package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.Open("./Go.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(file.Name())
}