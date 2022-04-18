package main // Package를 만들어서 사용할때 함수, 상수, 변수, 구조체의 이름이 대문자가 아니라면 외부 패키지로 노출되지 않는다.

import "example.com/hello"

func main(){
	hello.SayHello()
}
