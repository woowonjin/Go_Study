package main

import (
	"fmt"
	"net/http"
)

func main(){
	if response, err := http.Get("https://example.com"); err == nil {
		fmt.Println(response.StatusCode)
	}
}