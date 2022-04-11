package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, byte := range "Hello, Go!"{
			fmt.Printf("%c\n", byte)
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	fmt.Println("Goodbye")
}