package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	target := "127.0.0.1:30003"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}
	fmt.Println("Will send message now but will not care if address and port are available...")
	conn, err := net.Dial("udp", target)
	conn.Write([]byte("hello"))
	//	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Println("message sent")
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	//	n, dst, _ := conn.ReadFrom(p)
	conn.Close()
}
