package main

import (
	"fmt"
	"errors"
)

// func main() {
// 	defer func() {
// 		recover_var := recover()
// 		fmt.Println("var is : ", recover_var)
// 	}()
// 	panic("PANIC")
// }

func main() {
	defer func() {
		p := recover()
		if p == nil {
			return
		}
		err, ok := p.(error)
		if ok {
			fmt.Printf("%#v\n", err)
			return
		}
		panic(p)
	}()
	panic(errors.New("ERROR"))
}