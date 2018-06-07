package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var sources, targets []byte = readFiles("./data/sources.txt", "./data/targets.txt")

func random(min, max int) int {
	return rand.Intn(max-min) + min
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
