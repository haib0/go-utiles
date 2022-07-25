package service

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func getDir() []string {
	var file_list []string
	files, _ := ioutil.ReadDir(".")
	shellFmt := ".sh"
	if runtime.GOOS == "windows" {
		shellFmt = ".bat"
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), shellFmt) {
			file_list = append(file_list, f.Name())
		}
	}
	return file_list
}
