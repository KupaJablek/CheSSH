package util

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig()

	if err != nil {
		t.Fatalf(`ERROR: %s`, err.Error())
	}
}
