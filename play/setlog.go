package main

import (
	"log"
	"os"
)

func SetLog(logFile string) {
	_, err := os.Stat(logFile)
	if !os.IsExist(err) {
		os.Create(logFile)
	}

	fp, _ := os.Open(logFile)
	log.SetOutput(fp)
}
