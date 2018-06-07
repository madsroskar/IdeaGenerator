package main

import (
	"encoding/json"
	"fmt"
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
	err := bootstrap.ExecuteTemplate(w, p.templatePath, p.data)
	if err != nil {
		fmt.Println(err)
	}
}
