package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Page struct {
	Data         interface{}
	templatePath string
}

func (p Page) formatJsonData() interface{} {
	switch p.Data.(type) {
	case Idea:
		return map[string]string{"Idea": fmt.Sprintf("%s", p.Data.(Idea))}
	default:
		return p.Data
	}
}

func (p Page) renderJson(w http.ResponseWriter) {
	data := p.formatJsonData()
	json.NewEncoder(w).Encode(data)
}

func (p Page) renderHtml(w http.ResponseWriter) {
	err := bootstrap.ExecuteTemplate(w, p.templatePath, p)
	if err != nil {
		fmt.Println(err)
	}
}
