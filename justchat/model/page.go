package model

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename :="data/" + p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)
}

