package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/haib0/go-utiles/listentogether/server"
)

func Send() {
	for {
		sendTo()
	}
}

func sendTo() {
	cachedIPS := getCachedIP()
	cachedIP := ipFrom(cachedIPS)
	fmt.Printf("Send to (default is %s): ", cachedIPS)

	inputIPS, err := bufio.NewReader(os.Stdin).ReadString('\n')
	inputIPS = strings.TrimSpace(inputIPS)
	if err != nil {
		fmt.Println("err")
	}

	inputIP := ipFrom(inputIPS)

	if inputIP != nil {
		cacheIP(inputIPS)
		server.Server(inputIP)
	} else if cachedIP != nil {
		server.Server(cachedIP)
	}
}
