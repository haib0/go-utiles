package handler

import (
	"github.com/haib0/go-utiles/justtalk/utils"
	"net/http"
	"strings"
)

func pageHandlerPost(w http.ResponseWriter, r *http.Request) {
	pathList := strings.Split(r.URL.Path, "/")
	title := pathList[1]
	if title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		body := r.FormValue("body")
		if body == ""{
			http.Redirect(w, r, "/"+title+"/"+pathList[2], http.StatusFound)
			return
		}
		err := utils.CreateMessage(title, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(pathList) > 2 {
			http.Redirect(w, r, "/"+title+"/"+pathList[2], http.StatusFound)
			return
		}
		http.Redirect(w, r, "/"+title, http.StatusFound)
	}
}
