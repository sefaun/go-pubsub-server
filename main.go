package main

import (
	"bufio"
	"fmt"
	"go-pubsub-server/commands"
	"net"

	"github.com/jiyeyuran/go-eventemitter"
)

func Server(con net.Conn, em eventemitter.IEventEmitter) {

	client_content := NewClient()
	err := client_content.ClientContainer(con, em)

	if err != nil {
		fmt.Println(err)
	}
}

/*
Client Start
*/
type Client struct {
	ClientConnection *net.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) ClientContainer(client_connection net.Conn, em eventemitter.IEventEmitter) error {

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
			client.Subscribe(client_connection, result.Subscribe.Event, em)
		}

		if result.Publish.Cmd != "" {
			client.Publish(client_connection, result.Publish.Event, []byte(result.Publish.Message), em)
		}
	}

}

// func (client *Client) SendDataToClient(client_connection net.Conn, message string) {
// 	print("geldi")
// 	client_connection.Write([]byte(message))
// }

func (client *Client) Subscribe(client_connection net.Conn, event string, em eventemitter.IEventEmitter) {
	em.On(event, func(client_connection net.Conn, message string) {
		print("geldi")
		client_connection.Write([]byte(message))
	})
}

func (client *Client) Publish(client_connection net.Conn, event string, data []byte, em eventemitter.IEventEmitter) {
	em.Emit(event, client_connection, string(data))
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
	em := eventemitter.NewEventEmitter()

	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go Server(con, em)
	}
}
