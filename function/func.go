package main

import (
	"fmt"
	// "log"
)

// func sayHello(message string) (string, error) {
// 	return message, nil
// }

func sayHello() (helloGo string){
	helloGo = "Hello, Go!"
	return
}

func main(){
	// helloGo, err := sayHello("Hello, Go!")
	// if err != nil{
	// 	log.Fatal(err)
	// }
	helloGo := sayHello()
	fmt.Println(helloGo)
}


