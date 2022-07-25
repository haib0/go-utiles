package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/haib0/go-utiles/listentogether/io_api"
)

func Server(dst net.IP) {
	cpt := make(chan []byte, 1)

	cfg := io_api.StreamConfig{
		Channels:   2,
		SampleRate: 44100,
	}
	go io_api.Capture(context.Background(), cpt, cfg)

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	go sending()
	for {
		sendData := <-cpt
		// fmt.Println(sendData)
		_, err = listen.WriteToUDP(sendData, &net.UDPAddr{
			IP:   dst,
			Port: 30000,
		})
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}

}

func sending() {
	dots := []string{".  ", ".. ", "..."}
	for {
		for _, dot := range dots {
			fmt.Print("\rSending" + dot)
			time.Sleep(time.Second)
		}
	}
}
