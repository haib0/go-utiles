package main

import (
	"github.com/haib0/go-utiles/justchat/handler"
	"github.com/haib0/go-utiles/justchat/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", utils.MakeHandler(handler.ViewHandler))
	http.HandleFunc("/edit/", utils.MakeHandler(handler.EditHandler))
	http.HandleFunc("/save/", utils.MakeHandler(handler.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}