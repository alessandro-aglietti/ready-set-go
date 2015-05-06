package hello

import (
    "fmt"
    "net/http"
    "html/template"
    "math/rand"
	  "strconv"
    "time"
    "appengine"
    "appengine/datastore"
    "appengine/user"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/html-template", handlerHtmlTemplate)
    http.HandleFunc("/simple-form", handlerSimpleForm)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, mondo!")
}

/*
* NB: se il nome del campo non inizia per lettera maiuscola non è visibile all'esterno, ergo
* - non sarà scritto sul database
* - non sarà utilizzabile nei template
*/
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
    valoreCasuale := rand.Intn(300000000);
    data := BasicModel{"Sotto un valore casuale da 0 a 300000000", strconv.Itoa(valoreCasuale)}
  
    // unisco struttura dati e template html
    err = tmpl.Execute(w, data)
    if err != nil {
      // se avvengono errori nella creazione della pagina dinamica mi fermo
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
}

/*
* NB: se il nome del campo non inizia per lettera maiuscola non è visibile all'esterno, ergo
* - non sarà scritto sul database
* - non sarà utilizzabile nei template
*/
type Messaggio struct {
        Testo string
        Autore string
        Data time.Time
}
func (msg *Messaggio) DataRFC822() string {
  loc, _ := time.LoadLocation("Europe/Berlin")
  return msg.Data.In(loc).Format(time.RFC822)
}

func handlerSimpleForm(w http.ResponseWriter, r *http.Request) {
  
  c := appengine.NewContext(r)
  
  if ( r.Method == "GET" ) {
    c.Infof("handlerSimpleForm con metodo GET")
    // modalità visualizzazione, recuperare il template ed i dati
    // recupero il template
    tmpl, err := template.ParseFiles("templates/simple-form.html")
    if err != nil {
      // se avvengono errori nel recupero del template mi fermo
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    
    // recupero i dati salvati in precedenza
    q := datastore.NewQuery("Messaggio").Order("-Data").Limit(10)
    messaggi := make([]Messaggio, 0, 10)
    
    if _, err := q.GetAll(c, &messaggi); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  
    // unisco struttura dati e template html
    err = tmpl.Execute(w, messaggi)
    if err != nil {
      // se avvengono errori nella creazione della pagina dinamica mi fermo
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  } else if ( r.Method == "POST" ) {
    c.Infof("handlerSimpleForm con metodo POST")
    
    // inizio auth    
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }    
    // fine auth
    
    // modalità salvataggio dati
    messaggio := r.FormValue("messaggio")
    if messaggio != "" {
      // se il messaggio non è vuoto lo posso salvare!
      msg := Messaggio{ Testo: messaggio, Autore: u.Email, Data: time.Now() }
      
      // genero una nuova chiave incompleta che sarà completata durante l'inserimento
      key := datastore.NewIncompleteKey(c, "Messaggio", nil)
      
      // salvo il messaggio
      _, err := datastore.Put(c, key, &msg)
      if err != nil {
        // in caso di errori durante il salvataggio invio l'errore all'utente
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
    }
    
    // al termine delle operazioni rimando l'utente in GET sulla pagina
    http.Redirect(w, r, "/simple-form", http.StatusFound)
  }
}