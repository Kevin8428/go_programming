package main

import "net/http"


// DEFAULT MUX
func main() {
  http.HandleFunc("/", someFunc) // HandleFunc is in http library
  http.ListenAndServe(":8080", nil)
}

// NEW MUX, removed default mux
// ServeMux = HTTP request router = multiplexor = MUX
// compares incoming requests against a list of predefined URL paths,
// and calls the associated handler for the path whenver a match is found.
// func main() {
//   myMux := http.NewServeMux()
//   myMux.HandleFunc("/", someFunc) //someFunc is the handler
//   http.ListenAndServe(":8080", myMux)
// }

func someFunc(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("hello universe!"))
}
