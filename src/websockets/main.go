package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// construindo websocket http://www.pedromendes.com.br/2018/03/11/construindo-websockets-em-go/
// http://websocket.org/echo.html
var upgrader = websocket.Upgrader{
	// HandshakeTimeout:  1024,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// WriteBufferPool:   nil,
	// Subprotocols:      nil,
	//Error:             nil,
	CheckOrigin: func(r *http.Request) bool { return true },
	// EnableCompression: false,
}

func main() {
	http.HandleFunc("/", handler) // declaração path e o Handler
	http.ListenAndServe(":3000", nil)
	//fmt.Println("Hello!")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	for {

		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		// logando no console do webserver
		fmt.Println("Mensagem recebida: ", string(msg))

		// Devolvendo a mensagem recebida de volta para o cliente
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Fprintf(writer, "Hello")
}
