package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWebpageContent(urls []string) {
	for _, url := range urls {
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
	close(ch)
}

var ch chan string = make(chan string)

func main() {
	defer close(ch)

	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	go GetWebpageContent(urls)
	
	for b := range ch {
		fmt.Println(b)
	}
}