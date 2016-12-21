package main

import (
  "net/http"
)

type person struct { // lowercase == access only in this package: private
  fName string //fName is data
}

// with (p *person) ServeHTTP changes from function to method
// ServeHTTP is a method of any person struct
// instances if 'person' are created with ServeHTTP function as a method
// all instances will have fName data and ServeHTTP method
// p *person is called A RECEIVER
// ServeHTTP has receiver 'person'
// 'person' is receiving the method ServeHTTP
// since person implements handler interface, person becomes handler
// to implement handler interface, need ServeHTTP method with parameters
func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("First Name: " + p.fName))
}

func main() {

  personOne := &person{fName: "kevin"}
// since person implements handler interface, person becomes handler
// since 'person' is now a handler, it can be placed in ListenAndServe
  http.ListenAndServe(":8080", personOne)
}
