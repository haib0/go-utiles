package main

import (
	"github.com/haib0/go-utiles/justtalk/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		err = os.Mkdir("data", 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler.PageHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}