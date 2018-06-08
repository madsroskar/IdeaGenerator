package main

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func getRandomIdea(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("Content-Type")
	var page Page = Page{createRandomIdea(sources, targets), "idea"}
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

	idea, err := createIdea(sources, targets, source, target)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	var page Page = Page{idea, "idea"}
	if v == "application/json" {
		page.renderJson(w)
		return
	}
	page.renderHtml(w)
}