package handler

import (
	"github.com/haib0/go-utiles/justtalk/tmpl"
	"github.com/haib0/go-utiles/justtalk/utils"
	"net/http"
	"strconv"
	"strings"
)


func pageHandlerGet(w http.ResponseWriter, r *http.Request) {
	pathList := strings.Split(r.URL.Path, "/")
	title := pathList[1]
	if title == "" {
		_, _ = w.Write([]byte(tmpl.IndexTmpl))
		return
	}
	if len(pathList) < 3 {
		if title == "favicon.ico" {
			return
		}
		http.Redirect(w, r, "/"+title+"/0", http.StatusFound)
		return
	} else if len(pathList) > 3 {
		http.Redirect(w, r, "/"+title+"/"+pathList[2], http.StatusFound)
		return
	} else {
		i, err := strconv.ParseInt(pathList[2], 10, 0)
		if err != nil {
			http.Redirect(w, r, "/"+title+"/0", http.StatusFound)
			return
		}
		pageNo := int(i)
		if pageNo < 0 {
			http.Redirect(w, r, "/"+title+"/0", http.StatusFound)
		}
		page, err := utils.LoadPage(title, pageNo, 10)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if page.No > page.LastNo {
			http.Redirect(w, r, "/"+title+"/"+strconv.Itoa(page.LastNo), http.StatusFound)
		}
		utils.RenderTmpl(w, page)
	}
}
