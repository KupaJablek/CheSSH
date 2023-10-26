package online

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

/*
	NOT IMPLEMENTED FULLY
*/

func JoinSSHLobby(HOST string, PORT string, TYPE string) (*ssh.Client, error) {

	address := HOST + ":" + PORT

	var hostkeyCallback ssh.HostKeyCallback
	hostkeyCallback, err := knownhosts.New(CurrentUserHostsPath())
	if err != nil {
		errMsg := fmt.Sprint("HOST KEY ERROR: ", err.Error())
		return nil, errors.New(errMsg)
	}

	sshKeyPath := CurrentUserKeyPath()

	config := &ssh.ClientConfig{
		User:            "foo",
		HostKeyCallback: hostkeyCallback,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(PublicKeyFile(sshKeyPath)),
		},
	}

	conn, err := ssh.Dial(TYPE, address, config)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to dial: %s \tERROR: %s", address, err)
		return nil, errors.New(errMsg)
	}
	return conn, nil
}
