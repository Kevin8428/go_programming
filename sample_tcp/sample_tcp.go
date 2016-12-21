package main

import "net/http"

func main() {
  http.HandleFunc("/", someFunc) // HandleFunc is in http library
  http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("hello universe"))
}
