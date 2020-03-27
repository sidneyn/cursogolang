package main

import "net/url"

const (
	message2 = "Ping com socket gorilla"
	Charater = "\r\n\r\n"
)

// Client side
// shema - can be ws:// or wss://
//host, port - webSocket server

type  websocket struct {
	Scheme : Scheme
}
	var u := url.URL{
		Scheme : {"ws://"},
		Host : {"localhost"}: {"3000"},
		Path : "/"
	}
