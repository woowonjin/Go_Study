package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	resp, err := http.Get("http://www.naver.com")
// 	defer resp.Body.Close()
// 	if err != nil {
// 		panic(err)
// 	}

// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", string(data))
// }

func main() {
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}