package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWebpageContent(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- string(body)
}

var ch chan string = make(chan string)

func main() {
	defer close(ch)

	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go GetWebpageContent(url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}