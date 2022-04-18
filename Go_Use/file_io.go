package main

import (
	"io"
	"os"
	// "fmt"
)

func main() {
	fi, err := os.Open("temp/1.txt")
	defer fi.Close()
	if err != nil {
		panic(err)
	}

	fo, err := os.Create("temp/2.txt")
	defer fo.Close()
	if err != nil {
		panic(err)
	}

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// fmt.Println(buff, cnt)
		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}