package higor

import "testing"

func TestVersion(t *testing.T) {
	versionExpected := "v0.5.0"
	versionResult := Version

	if versionExpected != versionResult {
		t.Errorf("Version different. Expected: %s, Result: %s", versionExpected, versionResult)
	}
}

// TODO: higor.ReadCSV() implementation
