package main

import (
	"fmt"
	"go-pubsub-server/client"
	"net"

	"github.com/jiyeyuran/go-eventemitter"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	em := eventemitter.NewEventEmitter()

	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go Server(con, em)
	}
}

func Server(con net.Conn, em eventemitter.IEventEmitter) {
	client_content := client.Client{}

	go func() {
		err := client_content.ClientContainer(con, em)

		if err != nil {
			fmt.Println(err)
		}
	}()
}
