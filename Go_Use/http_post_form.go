package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
)

func main() {
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name" : {"Woo"}, "Age": {"10"}})
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)		
	}
}
