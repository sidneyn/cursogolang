package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type handlerSocket interface {
	handler2(port int)
}

const (
	Message       = "Pong"
	StopCharacter = "\r\n\r\n"
)

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	log.Println("lendo a listen")
	fmt.Println(listen)
	if err != nil {
		log.Fatalf("Socket listen port %d failed, %s", port, err)
		os.Exit(1)
	}

	defer listen.Close() // fecha no final socket

	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler2(conn)

	}

}
func handler2(conn net.Conn) {

	//	defer conn.Close() // fecha no final da conexão

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				break ILOOP
			}
		default:
			log.Fatalf("Receive data failde:%s", err)
			return
		}

	}
	w.Write([]byte(Message))
	w.Flush()
	log.Printf("Send: %s", Message)
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func main() {
	//port := 3333
	port2 := 3334

	//log.Print("SOCKET 1")
	//SocketServer(port)
	log.Print("SOCKET 2")
	SocketServer(port2)

}
