package handler

import (
	"github.com/haib0/go-utiles/justchat/utils"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := utils.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	utils.RenderTemplate(w, "view", p)
}
