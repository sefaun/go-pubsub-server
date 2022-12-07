package client

import (
	"bufio"
	"fmt"
	"go-pubsub-server/commands"
	"net"

	"github.com/jiyeyuran/go-eventemitter"
)

type Client struct {
	client_con net.Conn
	client_em  eventemitter.IEventEmitter
}

func (client *Client) NewClient(con net.Conn, em eventemitter.IEventEmitter) {
	client.client_con = con
	client.client_em = em

	go func() {
		err := client.ClientContainer()

		if err != nil {
			fmt.Println(err)
		}
	}()
}

func (client *Client) ClientContainer() error {
	defer client.client_con.Close()
	reader := bufio.NewReader(client.client_con)

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
			client.Subscribe(result.Subscribe.Event)
		}

		if result.Publish.Cmd != "" {
			fmt.Printf("cmd: %s - event: %s - message: %s", result.Publish.Cmd, result.Publish.Event, result.Publish.Message)
			client.Publish(result.Publish.Event, []byte(result.Publish.Message))
		}
	}

}

func (client *Client) SendDataToClient(message string) {
	client.client_con.Write([]byte(message))
}

func (client *Client) Subscribe(event string) {
	client.client_em.On(event, func(message string) {
		client.SendDataToClient(message)
	})
}

func (client *Client) Publish(event string, data []byte) {
	client.client_em.Emit(event, string(data))
}
