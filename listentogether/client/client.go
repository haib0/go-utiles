package client

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/haib0/go-utiles/listentogether/io_api"

	"github.com/gen2brain/malgo"
)

func Client() {
	pbk := make(chan []byte, 1)
	ctx := context.Background()
	cfg := io_api.StreamConfig{
		Format:     malgo.FormatF32,
		Channels:   2,
		SampleRate: 44100,
	}
	go io_api.Playback(ctx, pbk, cfg)

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	// RECV
	go recving()
	for {
		data := make([]byte, 4096)
		// n, remoteAddr, err := listen.ReadFromUDP(data)
		n, _, err := listen.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Err:", err)
			return
		}

		pbk <- data[:n]

		// fmt.Println("\nrecv: ", data[:n])
		// fmt.Println("\nFrom Addr: ", remoteAddr)
		// fmt.Println("\ncount: ", n)
	}

}

func recving() {
	dots := []string{".  ", ".. ", "..."}
	for {
		for _, dot := range dots {
			fmt.Print("\rReceiving" + dot)
			time.Sleep(time.Second)
		}
	}
}
