package utils

import (
	"io/ioutil"
	"github.com/haib0/go-utiles/justchat/model"
)

func LoadPage(title string) (*model.Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &model.Page{Title: title, Body: body}, nil
}
