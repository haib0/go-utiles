package main

import (
	"bufio"
	"fmt"
	"github.com/haib0/go-utiles/listentogether/app"
	"os"
	"strings"
)

func main() {
	fmt.Println("---Listen Together---")

	for {
		fmt.Println("[0] Send")
		fmt.Println("[1] Receive")
		fmt.Print("Choose: ")

		choose, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			continue
		}

		choose = strings.TrimSpace(choose)
		if choose == "0" {
			app.Send()
		} else if choose == "1" {
			app.Recv()
		} else if choose == "520" {
			app.Say520()
		}
		fmt.Println()
	}

}
