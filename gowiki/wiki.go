package main

// html/template allows us to keep HTML in separate file, allowing us to change the layout of our edit page without modifying the underlying Go code
import (
	"html/template"
	"io/ioutil"
  "net/http"
)

//this describes how data will be stored in memory
type Page struct {
    Title string
    Body  []byte // a byte slice, is a []byte rather than string because that is the type expected by the io library
}

// save method is built for persistent storage
// This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error
// This method will save the Page's Body to a text file. For simplicity, we will use the Title as the file name
// THIS SAVES PAGES:
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}


// loadPage constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values
// THIS LOADS PAGES
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename) //io.ReadFile returns returns []byte and error
    if err != nil { //could say if err == nil continue with results
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


// p2: The function then loads the page data, formats the page with a string of simple HTML, and writes it to w, the http.ResponseWriter.
// func main() {
//     p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//     p1.save()
//     p2, _ := loadPage("TestPage")
//     fmt.Println(string(p2.Body))
// }

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler) // loads teh page if exists and displays HTML form
	http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080", nil)
}









////////
