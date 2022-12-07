package utils_test

import (
	"go-pubsub-server/src/utils"
	"testing"
)

func TestContains(t *testing.T) {

	t.Run("Utils", func(t *testing.T) {
		result := utils.Contains([]string{"test_data"}, "test_data")

		if result == true {
			t.Log("Contains Test Successful")
		} else {
			t.Error("Contains Test Has Error")
		}

	})

}
