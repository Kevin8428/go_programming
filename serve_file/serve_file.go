package main

import (
  "net/http"
  "io/ioutil"
  "log"
)
// log is for debugging

type MyHandler struct{

}

// attaching ServeHTTP to MyHandler
// turns MyHandler into a handler because it implements Handler interface
// this says off request, get from URL the path and slice it from '/' so zero based index
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:] //get from URL path and remove slash
  log.Println(path)
  // path := "templates" + r.URL.Path

  // statically typed, once type is assigned it can't changes
  // compiler automatically throws in semicolons
  data, err := ioutil.ReadFile(string(path)) // read file at path from server and returns

  if err == nil {
    w.Write(data) //write to w from ServeHTTP method whatever data from ReadFile from URL
  } else {
    w.WriteHeader(404)
    w.Write([]byte("404 My Friend - " + http.StatusText(404)))
  }
}

func main() {
  http.Handle("/", new(MyHandler))
  http.ListenAndServe(":8080", nil)
}












////////
