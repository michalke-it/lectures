package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 30003
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Printf("Invalid port %v, err %v", os.Args[1], err)
			os.Exit(-1)
		} else {
			port = v
		}
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error while listening: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Listening at port %v...\n", addr.String())

	for {
		p := make([]byte, 1024)
		nn, raddr, err := server.ReadFromUDP(p)
		if err != nil {
			fmt.Printf("Read err %v", err)
			continue
		}

		msg := p[:nn]
		fmt.Printf("Received message [%s] from sender [%v]\n", msg, raddr)
	}
}
