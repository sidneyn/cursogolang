package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	message      = "Ping"
	StopCharater = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		fmt.Println("erro...........")
		os.Exit(1)

	}

	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharater))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {
	var (
		ip   = "127.0.0.1"
		port = 3333
	)
	SocketClient(ip, port)
}
