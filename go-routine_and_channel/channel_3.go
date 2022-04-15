package main

import (
	"fmt"
	"os"
	"time"
)

func Counter() {
	for i := 0;; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
}

func Stop() {
	time.Sleep(time.Second * 5)
	quit <- true
}

var ch chan int = make(chan int)
var quit chan bool = make(chan bool)


func main() {
	go Counter()
	go Stop()

	for {
		select {
		case count := <- ch:
			fmt.Println(count)
		case <-quit:
			os.Exit(1)
		}
	}
}