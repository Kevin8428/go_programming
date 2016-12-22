package main

import (
  "net/http"
  "text/template"
)

func main() {
  http.HandleFunc("/", myHandlerFunc)
  http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    tmpl.Execute(w, req.URL.Path[1:]) // passing in some data now
  }
}

const doc =`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>template</title>
  </head>
  <body>
    <h1>Hi {{.}}</h1>
  </body>
</html>
`


/*
{{.}}
the period is a 'pipeline'
pipeline: the path the template will take through the data to get to the data that needs to be injected
the period says 'use all the data'
test URLS:
http://localhost:8080/
http://localhost:8080/vermont
http://localhost:8080/boston/vermont
*/







//
