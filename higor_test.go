package higor

import "testing"

func TestPrintHelloHigor(t *testing.T) {
	value := PrintHelloHigor()
	if value != "Hello from higor" {
		t.Errorf("HellowHigor failed")
	}
}
