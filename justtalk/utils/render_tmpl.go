package utils

import (
	"html/template"
	"github.com/haib0/go-utiles/justtalk/model"
	"github.com/haib0/go-utiles/justtalk/tmpl"
	"log"
	"net/http"
)

func RenderTmpl(w http.ResponseWriter, page *model.Page)  {
	t, err := template.New("page").Parse(tmpl.PageTmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "page", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
