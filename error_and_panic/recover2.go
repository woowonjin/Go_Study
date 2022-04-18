package main

import (
	"os"
	"fmt"
)

func main() {
	openFile("Invalid.txt")

	// recover에 의해 아래 코드는 실행됨
	println("Done")
}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}