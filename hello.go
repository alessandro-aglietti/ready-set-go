package hello

import (
    "fmt"
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/html-template", handlerHtmlTemplate)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

type Model struct {
	Titolo1, Titolo2 string
}
func handlerHtmlTemplate(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/basic.html")
    
    err = tmpl.Execute(w, Model{"Titolo 1", "Titolo 2"})
  
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

