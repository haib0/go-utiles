package main

import (
	"flag"

	"github.com/haib0/go-utiles/ylgyproxy/proxy"
)

func main() {
	port := flag.Int("port", 10806, "port of justproxy")
	flag.Parse()
	proxy.RunProxy(*port)
}
