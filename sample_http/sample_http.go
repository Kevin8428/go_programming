// hello folder is a command
// hello.go is a command source

package main

import ("net/http" ; `io`)

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
    `<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello World! Hello path
  </body>
</html>`,
  )
}

func contact (res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello World! Contact Path
  </body>
</html>`,
	)
}

func main() {
  http.HandleFunc("/hello", hello)
	http.HandleFunc("/contact", contact)
  http.ListenAndServe(":9000", nil)
}
