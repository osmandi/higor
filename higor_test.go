package higor

import (
	"testing"
)

func TestHelloHigor(t *testing.T) {

	resultMessage := HelloHigor()
	expectedMessage := "Hello from Higor :) v0.2.1"

	if resultMessage != expectedMessage {
		t.Errorf("Message expected: '%s' but received: '%s'", expectedMessage, resultMessage)
	}

}
