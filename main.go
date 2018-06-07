package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var sources, targets []byte = readFiles("./data/sources.txt", "./data/targets.txt")

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomIdea(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("Content-Type")
	var page Page = Page{createRandomIdea(sources, targets), "./tmpl/idea.html"}
	if v == "application/json" {
		page.renderJson(w)
		return
	}
	page.renderHtml(w)
}

func getIdea(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("Content-Type")
	vars := mux.Vars(r)

	source, serr := strconv.Atoi(vars["source"])
	target, terr := strconv.Atoi(vars["target"])

	if serr != nil || terr != nil {
		return
	}

	var page Page = Page{createIdea(sources, targets, source, target), "./tmpl/idea.html"}
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
	router.HandleFunc("/", getRandomIdea).Methods("GET")
	router.HandleFunc("/{source:[0-9]+}/{target:[0-9]+}", getIdea).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
