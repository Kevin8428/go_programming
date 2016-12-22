package main


// text/template doesn't do any escaping
// html/template is a wrapper of text/template with all same functionality
// allows for escaping and other features
// html/template is therefore safer and won't run alert script in this example
import (
  "net/http"
  "text/template"
)

// second Message works with <body><script>{{.}}</script></body> and "text/template"
// third Message works with <body><p>{{.}}</p></body> and "text/template"

// var Message string = "more beer, please"
// var Message string = "alert('alert: you have been pwned')"
var Message string = "<script>alert('script alert')</script>"

func main() {
  http.HandleFunc("/", myHandlerFunc)
  http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  tmpl, err := template.New("anyNameForTemplate").Parse(doc)
  if err == nil {
    tmpl.Execute(w, Message)
  }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
  <meta charset="UTF-8">
  <title>Injection Safe</title>
</head>
<body>
  <p>{{.}}</p>
</body>
</html>
`






//
