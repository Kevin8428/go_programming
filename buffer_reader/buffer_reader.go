package main

import (
  "net/http"
  "strings"
  "log"
  "os"
  "bufio"
)

type MyHandler struct {

}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  log.Println(path)
  f, err := os.Open(path)

  if err == nil {
    bufferReader := bufio.NewReader(f)

    var contentType string

    if strings.HasSuffix(path, ".css") {
      log.Println("aaa")
      contentType = "text/css"
    } else if strings.HasSuffix(path, ".html") {
      log.Println("bbb")
      contentType = "text/html"
    } else if strings.HasSuffix(path, ".js"){
      contentType = "application/javascript"
    } else if strings.HasSuffix(path, ".png"){
      contentType = "image/png"
    } else if strings.HasSuffix(path, ".svg"){
      contentType = "image/svg+xml"
    } else {
      contentType = "text/plain"
    }

    w.Header().Add("Content-Type", contentType)
    w.Write(data)
  } else {
    w.WriteHeader(404)
    w.Write([]byte("404 my friend - " + http.StatusText(404)))
  }
}

func main() {
  http.Handle("/", new(MyHandler))
  http.ListenAndServe(":8080", nil)
}





//
