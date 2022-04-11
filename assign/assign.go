package main

import (
	"fmt"
	"reflect"
)

var message string

var hi string = "Hi"
var whoAreYou = "Who are You?"

// goodBye := "GoodBye" -> Error
/*
함수 외부에서는 모든 문장이 var, func와 같은 키워드로 시작하므로 키워드로 시작하지 않는
단축 선언은 함수 외부에서 사용할 수 없다.
*/

var (
	temp1 string
	temp2 string
	temp3 string
) // 여러 변수 선언

func main(){
	message = "Hello, Go!"
	fmt.Println(reflect.TypeOf(hi))
	// a, b, c := "test", 1, 0.5
	// fmt.Printf("%#s, %#d, %#f", &a, b, c)
	// string := 10
	// fmt.Println(string)
	var message string
	message = "is good?"
	fmt.Println(message)
}