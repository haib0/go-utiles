package checkssh

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func CheckTCP(addr string) (bool, error) {
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		return false, fmt.Errorf("TCP dial %s timeout: %s", addr, err)
	}
	defer conn.Close()
	return true, nil
}

func CheckSSH(addr, user, passwd string) (bool, error) {
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return false, fmt.Errorf("SSH %s error: %s", addr, err)
	} else {
		client.Close()
		return true, nil
	}
}

func Check(addr, user, passwd string, wg sync.WaitGroup) {
	tcpOK, _ := CheckTCP(addr)
	if tcpOK {
		sshOK, _ := CheckSSH(addr, user, passwd)
		if sshOK {
			log.Println("SSH Success: ", addr)
		}
	}

	wg.Done()
}
