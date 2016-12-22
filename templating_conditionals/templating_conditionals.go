package main

import (
  "net/http"
  "text/template"
)

type Context struct {
    FirstName string
    Message string
    URL string
}


func main() {
  http.HandleFunc("/", myHandlerFunc) // forward slash catches all except direct match to /nobeer
  http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{"friend", "dynamic templating FTW", req.URL.Path} // stores data in instance of Context struct
    tmpl.Execute(w, context)
  }
}


// BELOW: eq is operator and operands follow
// .URL is saying variable URL
const doc =`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>template</title>
  </head>
  <body>
    {{if eq .URL "/nobeer"}}
      <h1>we're out of beer, sorry<h1>
      <h1>still, {{.Message}}<h1>
    {{else}}
      <h1>sure, grab another beer, {{.FirstName}}<h1>
      <h1>{{.Message}}<h1>
    {{end}}
    <hr>
    <h2>here's all the data:</h2>
    <p>{{.}}</p>
  </body>
</html>
`








///
