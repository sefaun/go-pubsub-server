package commands_test

import (
	"go-pubsub-server/src/commands"
	"testing"
)

func TestFindCommand(t *testing.T) {

	t.Run("Publish Commands", func(t *testing.T) {
		_, err := commands.FindCommand("PUB:Event:Data")

		if err != nil {
			t.Error("FindCommand Test Has Error")
		}
	})

	t.Run("Subscribe Commands", func(t *testing.T) {
		_, err := commands.FindCommand("SUB:Event")

		if err != nil {
			t.Error("FindCommand Test Has Error")
		}
	})

	t.Run("FetchSubscribeCommand Commands", func(t *testing.T) {
		result, err := commands.FetchSubscribeCommand([]string{"SUB", "SubEvent"})

		if err != nil {
			t.Error("FetchSubscribeCommand Test Has Error")
		}

		if result.Cmd != "SUB" {
			t.Error("SUB Command Test Has Error")
		}

		if result.Event != "SubEvent" {
			t.Error("SubEvent Command Test Has Error")
		}
	})

	t.Run("FetchPublishCommand Commands", func(t *testing.T) {
		result, err := commands.FetchPublishCommand([]string{"PUB", "PubEvent", "Message"})

		if err != nil {
			t.Error("FetchPublishCommand Test Has Error")
		}

		if result.Cmd != "PUB" {
			t.Error("PUB Command Test Has Error")
		}

		if result.Event != "PubEvent" {
			t.Error("PubEvent Command Test Has Error")
		}

		if result.Message != "Message" {
			t.Error("Message Command Test Has Error")
		}
	})

}
