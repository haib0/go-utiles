package utils

import (
	"html/template"
	"github.com/haib0/go-utiles/justchat/model"
	"net/http"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
