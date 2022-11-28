package commands

import (
	"errors"
	"strings"
)

const (
	Subscribe = "SUB"
	Publish   = "PUB"
)

type SubscribeResponse struct {
	Cmd   string `json:"cmd"`
	Event string `json:"event"`
}

type PublishResponse struct {
	Cmd     string `json:"cmd"`
	Event   string `json:"event"`
	Message string `json:"message"`
}

type Commands struct {
	Subscribe SubscribeResponse
	Publish   PublishResponse
}

//data -> SUB:Event
func FetchSubscribeCommand(data []string) (SubscribeResponse, error) {
	if len(data) != 2 {
		return SubscribeResponse{}, errors.New("subscribe protocol error")
	}

	return SubscribeResponse{
		Cmd:   Subscribe,
		Event: data[1],
	}, nil
}

//data -> PUB:Event:Data
func FetchPublishCommand(data []string) (PublishResponse, error) {
	if len(data) != 3 {
		return PublishResponse{}, errors.New("publish protocol error")
	}

	return PublishResponse{
		Cmd:     Publish,
		Event:   data[1],
		Message: data[2],
	}, nil
}

func FindCommand(messages string) (Commands, error) {
	command_messages := strings.Split(messages, ":")

	if len(command_messages) == 0 {
		return Commands{}, errors.New("protocol error")
	}

	switch command_messages[0] {
	case Subscribe:
		result, err := FetchSubscribeCommand(command_messages)

		if err != nil {
			return Commands{}, err
		}

		return Commands{
			Subscribe: result,
		}, nil

	case Publish:
		result, err := FetchPublishCommand(command_messages)

		if err != nil {
			return Commands{}, err
		}

		return Commands{
			Publish: result,
		}, nil

	default:
		return Commands{}, errors.New("protocol error")
	}

}
