package main

import (
	"path/filepath"
	"os"
	"fmt"
	"math/rand"
	"time"
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "html/template"
)

var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func getIdea(w http.ResponseWriter, r *http.Request) {
	//i := make(map[string]string)
	//i["Idea"] = fmt.Sprintf("%s", createIdea(sources, targets))
	//json.NewEncoder(w).Encode(i)


    t, err := template.ParseFiles("./tmpl/idea.html")  // Parse template file.
    if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	idea := createIdea(sources, targets) // Get current user infomration.
    t.Execute(w, idea)  // merge.
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

