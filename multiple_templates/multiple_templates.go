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
  templates := template.New("template")
  templates.New("test").Parse(doc)
  templates.New("header").Parse(head)
  templates.New("footer").Parse(foot)
    context := Context{
      "kevin",
      "I'd like a beer please",
      req.URL.Path,
      []string{"Alchemist", "switchback", "lawsons finest"},
      "Favorite Beers",
      }
    templates.Lookup("test").Execute(w, context)
}

// {{template "header" .Title}} passes into header file just data .Title
// in const header therefore, the new pipeline is only the value of .Title
const doc = `
{{template "header" .Title}}

<body>

  <h1>{{.FirstName}} says, "{{.Message}}"</h1>

  {{if eq .URL "/nobeer"}}
    <h2>we're out of beer, sorry {{.FirstName}}</h2>
  {{else}}
    <h2>sure, grab another, {{.FirstName}}</h2>
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
{{template "footer"}}
`

const head = `
<!DOCTYPE html>
<html>
<head lang="en">
  <meta charset="UTF-8">
  <title>{{.}}</title>
</head>
`
const foot = `
<hr>
<footer>footer text</footer>
</html>
`





///
