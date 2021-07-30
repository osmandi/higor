package higor

import (
	"os/exec"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	c, b := exec.Command("git", "branch", "--show-current"), new(strings.Builder)
	c.Stdout = b
	c.Run()
	versionExpected := strings.TrimRight(b.String(), "\n")
	versionResult := Version

	if versionExpected != versionResult {
		t.Errorf("Version different. Expected: %s, Result: %s", versionExpected, versionResult)
	}
}

// TODO: higor.ReadCSV() implementation
