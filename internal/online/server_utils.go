package online

import (
	"io/ioutil"
	"os/user"
	"runtime"

	"golang.org/x/crypto/ssh"
)

func PublicKeyFile(file string) ssh.Signer {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return key
}

func CurrentUserKeyPath() string {
	platform := runtime.GOOS
	if platform == "Windows" {
		// not yet implemented
		return ""
	} else {
		user, _ := user.Current()
		return "/home/" + user.Username + "/.ssh/id_rsa"
	}
}

func CurrentUserHostsPath() string {
	platform := runtime.GOOS
	if platform == "Windows" {
		// not yet implemented
		return ""
	} else {
		user, _ := user.Current()
		return "/home/" + user.Username + "/.ssh/known_hosts"
	}
}
