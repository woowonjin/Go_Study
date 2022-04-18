package main

import (
	"os"
)

func main() {
	defer println("OK")
	openFile("Invalid.txt")
	// defer println("OK")
	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}