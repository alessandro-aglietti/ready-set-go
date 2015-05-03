package hello

import (
    "fmt"
    "net/http"
    "html/template"
    "math/rand"
	  "strconv"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/html-template", handlerHtmlTemplate)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

type BasicModel struct {
	Titolo1, Titolo2 string
}
func handlerHtmlTemplate(w http.ResponseWriter, r *http.Request) {
  
    // recupero il template
    tmpl, err := template.ParseFiles("templates/basic.html")
    if err != nil {
      // se avvengono errori nel recupero del template mi fermo
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    
    // creo la struttura dati composta da un titolo fisso ed un valore casuale
    valoreCasuale := rand.Intn(3000);
    data := BasicModel{"Sotto un valore casuale da 0 a 3000", strconv.Itoa(valoreCasuale)}
  
    // unisco struttura dati e template html
    err = tmpl.Execute(w, data)
    if err != nil {
      // se avvengono errori nella creazione della pagina dinamica mi fermo
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}