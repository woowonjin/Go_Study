package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	defer resp.Body.Close()
	if err != nil{
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
}