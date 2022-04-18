package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{"wonjin", 10}
	pbytes, _ := json.Marshal(person)
	fmt.Println(pbytes)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
}