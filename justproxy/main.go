package main

import "flag"

func main() {
	port := flag.Int("port", 10806, "port of justproxy")
	flag.Parse()
	RunProxy(*port)
}
