package tests

import (
	"fmt"
	"testing"

	hg "github.com/osmandi/higor"
)

func TestHelloHigor(t *testing.T) {

	resultMessage := hg.HelloHigor()
	expectedMessage := fmt.Sprintf("Hello from Higor :), %s", hg.Version)

	if resultMessage != expectedMessage {
		t.Errorf("Message expected: '%s' but received: '%s'", expectedMessage, resultMessage)
	}

}
