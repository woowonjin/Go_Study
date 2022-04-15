package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func GetWebpageContent(url string) string {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go GetWebpageContent("https://example.com")
	go GetWebpageContent("https://golang.org")
	go GetWebpageContent("https://golang.org/doc")
	wg.Wait()
}