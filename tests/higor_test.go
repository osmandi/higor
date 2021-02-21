package tests

import (
	"testing"

	hg "github.com/osmandi/higor"
)

func TestHelloHigor(t *testing.T) {

	resultMessage := hg.HelloHigor()
	expectedMessage := "Hello from Higor :) v0.2.0"

	if resultMessage != expectedMessage {
		t.Errorf("Message expected: '%s' but received: '%s'", expectedMessage, resultMessage)
	}

}
