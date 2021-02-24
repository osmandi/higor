package higor

import (
	"fmt"
)

var Version string = "v0.2.1"

// HelloHigor Print a simple message to check if Higor are installed correctly
// and print the version installed
func HelloHigor() string {

	helloMessage := fmt.Sprintf("Hello from Higor :) %s", Version)
	return helloMessage
}
