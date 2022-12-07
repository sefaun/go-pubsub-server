package main

import (
	"go-pubsub-server/src/client"
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

		new_client := client.Client{}
		go new_client.NewClient(con, em)
	}
}
