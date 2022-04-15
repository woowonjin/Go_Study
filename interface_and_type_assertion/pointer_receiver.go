package main

import "fmt"

type Develop interface {
	Coding()
}

type Developer struct {
	Languages []string
}

func (d *Developer) Coding() {
	for _, l := range d.Languages {
		fmt.Println(l)
	}
}

func main() {
	var d Develop = &Developer{[]string{"Go", "PHP", "Javascript"}}
	d.Coding()

	// d.Coding()
}