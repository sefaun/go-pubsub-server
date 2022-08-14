package main

import (
	"bufio"
	"fmt"
	"go-pubsub-server/commands"
	"net"
)

/*
	EventEmitter Start
*/
type EventEmitter map[string]chan string

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{}
}

func (event_emitter EventEmitter) Server(con net.Conn) {

	client_content := NewClient()
	client_content.ClientContainer(con, event_emitter)

}

/*
	EventEmitter End
*/

/*
Client Start
*/
type Client struct {
	ClientConnection *net.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) ClientContainer(client_connection net.Conn, event_emitter EventEmitter) error {

	defer client_connection.Close()
	reader := bufio.NewReader(client_connection)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			return err
		}
		result, err := commands.FindCommand(message)

		if err != nil {
			return err
		}

		if result.Subscribe.Cmd != "" {
			client.Subscribe(client_connection, result.Subscribe.Event, event_emitter)
		}

		if result.Publish.Cmd != "" {
			client.Publish(result.Publish.Event, []byte(result.Publish.Message), event_emitter)
		}
	}

}

func (client *Client) Subscribe(client_connection net.Conn, event string, event_emitter EventEmitter) {
	event_emitter[event] = make(chan string)

	go func() {
		for {
			select {
			case message := <-event_emitter[event]:
				fmt.Println(message)
				client_connection.Write([]byte(message))
			}
		}
	}()

}

func (client *Client) Publish(event string, data []byte, event_emitter EventEmitter) {
	event_emitter[event] <- string(data)
}

/*
	Client End
*/

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	event_emitter := NewEventEmitter()

	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go event_emitter.Server(con)
	}
}
