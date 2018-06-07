package main

import (
	"path/filepath"
	"os"
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "html/template"
)

var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

type Page struct {
	data interface{}
	templatePath string
}

func (p Page) formatJsonData() interface{}  {
	switch p.data.(type) {
		case Idea:
			return map[string]string{ "Idea": fmt.Sprintf("%s", p.data.(Idea)) }
		default:
			return p
	}
}

func (p Page) renderJson(w http.ResponseWriter) {
	data := p.formatJsonData();
	json.NewEncoder(w).Encode(data)
}

func (p Page) renderHtml(w http.ResponseWriter) {
    t, err := template.ParseFiles(p.templatePath)
    if err != nil {
		fmt.Println(err)
	}
	
	t.Execute(w, p.data)
}

func getIdea(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("Content-Type")	
	var page Page = Page{createIdea(sources, targets), "./tmpl/idea.html"}
	if v == "application/json" {
		page.renderJson(w)
		return
	}
 	page.renderHtml(w)
}

func main() {
	rand.Seed(time.Now().Unix())
    
	_, err := os.Stat(filepath.Join(".", "tmpl", "css", "style.css"))

	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./tmpl/css/"))))
	router.HandleFunc("/", getIdea).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

