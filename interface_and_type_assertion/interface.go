package main

import "fmt"

type Develop interface {
	Coding()
}

type Developer struct {
	Languages []string
}

func Work(developer Develop) {
	developer.Coding()
	// developer.OtherMethod() -> Error
	if developer, ok := developer.(Developer); ok {
		developer.OtherMethod()
	}
}

func (d Developer) Coding() {
	for _, l := range d.Languages {
		fmt.Println(l)
	}
}

func (d Developer) OtherMethod() {
	fmt.Println("Other Method!!")
}

func main() {
	d := Developer{[]string{"Go", "PHP", "Javascript"}}
	Work(d)

	// d.Coding()
}