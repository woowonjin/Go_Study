package main

import "fmt"

func each(arr []string, iterator func(string)){
	for _, v := range arr {
		iterator(v)
	}
}

func main(){
	p := func(v string){
		fmt.Println(v)
	}
	each([]string{"Hello, Go!", "Who are you?"}, p)
}