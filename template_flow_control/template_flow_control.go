package main

import (
  "net/http"
  "text/template"
)

// Beers is a slice of strings, slice is a wrapper around an array
type Context struct {
    FirstName string
    Message string
    URL string
    Beers []string
    Title string
}


func main() {
  http.HandleFunc("/", myHandlerFunc) // forward slash catches all except direct match to /nobeer
  http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    context := Context{
      "kevin",
      "I'd like a beer please",
      req.URL.Path,
      []string{"Alchemist", "switchback", "lawsons finest"},
      "Favorite Beers",
      }
    tmpl.Execute(w, context)
  }
}


// now doing loops
const doc =`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.FirstName}} says, "{{.Message}}"</h1>
    {{if eq .URL "/nobeer"}}
      <h1>we're out of beer, {{.FirstName}} sorry<h1>
    {{else}}
      <h1>sure, grab another beer, {{.FirstName}}<h1>
      <ul>
        {{range .Beers}}
        <li>{{.}}</li>
        {{end}}
      </ul>
    {{end}}
    <hr>
    <h2>here's all the data:</h2>
    <p>{{.}}</p>
  </body>
</html>
`








///
