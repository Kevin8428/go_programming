package main

import (
  "net/http"
  "text/template"
)

type Context struct {
    FirstName string
    Message string
}


const doc =`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>template</title>
  </head>
  <body>
    <h1>Hi {{.FirstName}}</h1>
    <p>{{.Message}}</p>
  </body>
</html>
`

func main() {
  http.HandleFunc("/boston", bostonFunc)
  http.HandleFunc("/vermont", vermontFunc)
  http.HandleFunc("/maine", maineFunc)
  http.HandleFunc("/", rootFunc)
  http.ListenAndServe(":8080", nil)
}

func bostonFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{"boston", "love that dirty water"}
    tmpl.Execute(w, context) // passing in some data now
  }
}

func vermontFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{"vermont", "go cats!"}
    tmpl.Execute(w, context) // passing in some data now
  }
}

func maineFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{"maine", "go blackbears!"}
    tmpl.Execute(w, context) // passing in some data now
  }
}

func rootFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{"index", "dynamic templating FTW"}
    tmpl.Execute(w, context) // passing in some data now
  }
}







///
