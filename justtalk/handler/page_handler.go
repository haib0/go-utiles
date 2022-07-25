package handler

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pageHandlerGet(w, r)
	} else if r.Method == "POST" {
		pageHandlerPost(w, r)
	} else {
		http.NotFound(w, r)
	}
}
