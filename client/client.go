package client

import (
	"bufio"
	"fmt"
	"go-pubsub-server/commands"
	"net"

	"github.com/jiyeyuran/go-eventemitter"
)

type Client struct{}

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
			fmt.Printf("cmd: %s - event: %s", result.Subscribe.Cmd, result.Subscribe.Event)
			client.Subscribe(client_connection, result.Subscribe.Event, em)
		}

		if result.Publish.Cmd != "" {
			fmt.Printf("cmd: %s - event: %s - message: %s", result.Publish.Cmd, result.Publish.Event, result.Publish.Message)
			client.Publish(client_connection, result.Publish.Event, []byte(result.Publish.Message), em)
		}
	}

}

func (client *Client) SendDataToClient(client_connection net.Conn, message string) {
	print("sending data")
	client_connection.Write([]byte(message))
}

func (client *Client) Subscribe(client_connection net.Conn, event string, em eventemitter.IEventEmitter) {
	em.On(event, func(client_connection net.Conn, message string) {
		client.SendDataToClient(client_connection, message)
	})
}

func (client *Client) Publish(client_connection net.Conn, event string, data []byte, em eventemitter.IEventEmitter) {
	em.Emit(event, client_connection, string(data))
}
