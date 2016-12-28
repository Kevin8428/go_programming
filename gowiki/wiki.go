// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//   "net/http"
//   "log"
// )
//
// //this describes how data will be stored in memory
// type Page struct {
//     Title string
//     Body  []byte // a byte slice, is a []byte rather than string because that is the type expected by the io library
// }
//
// // save method is built for persistent storage
// // This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error
// // This method will save the Page's Body to a text file. For simplicity, we will use the Title as the file name
// // THIS SAVES PAGES:
// func (p *Page) save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }
//
//
// // loadPage constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values
// // THIS LOADS PAGES
// func loadPage(title string) (*Page, error) {
//     filename := title + ".txt"
//     body, err := ioutil.ReadFile(filename) //io.ReadFile returns returns []byte and error
//     if err != nil { //could say if err == nil continue with results
//         return nil, err
//     }
//     return &Page{Title: title, Body: body}, nil
// }
//
// func viewHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/view/"):]
//     log.Println(title)
//     p, _ := loadPage(title) // _ ignores the error return value from loadPage
//     fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }
//
// // p2: The function then loads the page data, formats the page with a string of simple HTML, and writes it to w, the http.ResponseWriter.
// // func main() {
// //     p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
// //     p1.save()
// //     p2, _ := loadPage("TestPage")
// //     fmt.Println(string(p2.Body))
// // }
//
// func main() {
//     http.HandleFunc("/view/", viewHandler)
//     http.ListenAndServe(":8080", nil)
// }
//
//
//
//
//
//
//
//
//
// ////////




package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
