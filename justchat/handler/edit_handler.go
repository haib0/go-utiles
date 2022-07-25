package handler

import (
	"github.com/haib0/go-utiles/justchat/model"
	"github.com/haib0/go-utiles/justchat/utils"
	"net/http"
)

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := utils.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	utils.RenderTemplate(w, "edit", p)
}