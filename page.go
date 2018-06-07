package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	data         interface{}
	templatePath string
}

func (p Page) formatJsonData() interface{} {
	switch p.data.(type) {
	case Idea:
		return map[string]string{"Idea": fmt.Sprintf("%s", p.data.(Idea))}
	default:
		return p
	}
}

func (p Page) renderJson(w http.ResponseWriter) {
	data := p.formatJsonData()
	json.NewEncoder(w).Encode(data)
}

func (p Page) renderHtml(w http.ResponseWriter) {
	t, err := template.ParseFiles(p.templatePath)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, p.data)
}
