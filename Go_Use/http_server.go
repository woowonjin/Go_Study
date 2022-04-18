package main

import (
	"net/http"
)

// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
// 		w.Write([]byte("Hello World"))
// 	})

// 	http.ListenAndServe(":8000", nil)
// }

func main() {
	http.Handle("/", new(testHandler))
	http.ListenAndServe(":8000", nil)
}

type testHandler struct {
	http.Handler // Embedding
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}