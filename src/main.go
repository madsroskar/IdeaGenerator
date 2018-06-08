package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var sources, targets []byte = readFiles(os.Getenv("SOURCES_FILE"), os.Getenv("TARGETS_FILE"))
var bootstrap *template.Template

func random(min, max int) int {
	return rand.Intn(max-min) + min
}



func main() {
	rand.Seed(time.Now().Unix())

	_, err := os.Stat(filepath.Join(".", "tmpl", "css", "style.css"))
	bootstrap, err = template.ParseGlob(os.Getenv("TMPL_DIR") + "*.html")

	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(os.Getenv("TMPL_DIR") + "css/"))))
	router.HandleFunc("/", getRandomIdea).Methods("GET")
	router.HandleFunc("/{source:[0-9]+}/{target:[0-9]+}", getIdea).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
