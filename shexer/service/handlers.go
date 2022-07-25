package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func StartService() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/run/", runHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").Parse(indexTmpl)
	if err != nil {
		log.Fatal(err)
	}
	file_list := getDir()
	err = t.ExecuteTemplate(w, "index", file_list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, "/")
	shell := path[len(path)-1]
	fmt.Fprintf(w, "running %s\n\n", shell)
	cmd := exec.Command("echo")
	switch runtime.GOOS {
	case "windows":
		{
			cmd = exec.Command(shell)
		}
	case "linux":
		{
			cmd = exec.Command("sh", shell)
		}
	}
	cmd.Stdout = w
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "\n\nrun error: %s", err.Error())
	}
}
