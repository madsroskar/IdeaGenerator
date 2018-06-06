package main

import (
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func getIdea(w http.ResponseWriter, r *http.Request) {
	i := make(map[string]string)
	i["Idea"] = fmt.Sprintf("%s", createIdea(sources, targets))
	json.NewEncoder(w).Encode(i)
}

func main() {
	rand.Seed(time.Now().Unix())

	router := mux.NewRouter()
	router.HandleFunc("/", getIdea).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

