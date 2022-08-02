package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/elazarl/goproxy"
)

func RunProxy(port int) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	addr := ":" + strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
