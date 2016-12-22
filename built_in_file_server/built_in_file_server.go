package main

import "net/http"


// anonymous function, handler is http.FileServer
func main() {
  http.ListenAndServe(":8080", http.FileServer(http.Dir("")))
}
