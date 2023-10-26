package online

import (
	"os"
	"os/user"
	"runtime"
	"testing"
)

func TestGetDefaultPublicKey(t *testing.T) {
	user, _ := user.Current()
	platform := runtime.GOOS
	if platform == "Windows" {
		//TODO: test windows default path
	} else {
		filePath := "/home/" + user.Username + "/.ssh/id_rsa"
		key := PublicKeyFile(filePath)
		if key == nil {
			t.Fatalf(`PublicKeyFile(%s) = nil. Failed to access key`, filePath)
		}
	}
}

func TestGetDefaultKnownHosts(t *testing.T) {

	user, _ := user.Current()
	platform := runtime.GOOS
	if platform == "Windows" {
		//TODO: test windows default path
	} else {
		path := "/home/" + user.Username + "/.ssh/known_hosts"
		if _, err := os.Stat(path); err != nil {
			t.Fatalf(`Known hosts not found at: %s`, path)
		}
	}
}
