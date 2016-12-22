/*
templates:
- New
- Parse
- Execute
*/

package main

import (
  "net/http"
  "text/template"
)

func main() {
  http.HandleFunc("/", myHandlerFunc)
  http.ListenAndServe(":8080", nil) // nil means use default Mux
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
  w.Header().Add("Content-Type", "text/html")
  // create a New template,
  tmpl, err := template.New("anyNameForTemplate").Parse(doc) // later will grab from file, for ease placed below
  if err == nil { // execute tmpl if no error
    tmpl.Execute(w, nil) // 'w' - write it back to response, nil means no data to pass in
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
    <h1>template</h1>
  </body>
</html>
`

// templates are: create a new template, parse it and execute it.


/*
template.New
basically create a template "object"
go doesn't have constructors like in OO languages
go is type based
go's struct type often used like an object
-- it holds data
-- it can have methods (the struct must be a "receiver" on a function)
if any initialization is needed, then create a method that does this and return the instance:
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template/#New
*/


/*
Template.Parse
basically put your template into the template "object"
func (t*Template) parse(text string) (*Ttemplate, error)
parse parses a string into a template. Nested template destination will be associated with the top-level template t. Parse may be called multiple times to parse definitions of templates to associate with t
*/

/*
template.Execute
basically merge your template with data
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
Execute applies a parsed template to the specified data object, and writes the output to wr.
If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer.
A template may be executed safely in parallel
*/



//
