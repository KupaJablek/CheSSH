package online

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
)

/*
	NOT IMPLEMENTED FULLY
*/

func CreateSSHServer(HOST string, PORT string, TYPE string) (net.Conn, error) {
	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "foo" && string(pass) == "bar" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	filePath := CurrentUserKeyPath()
	key, err := ioutil.ReadFile(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("unable to read private key: %v", err)
		return nil, errors.New(errMsg)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		errMsg := fmt.Sprintf("unable to parse private key: %v", err)
		return nil, errors.New(errMsg)
	}
	config.AddHostKey(signer)

	address := HOST + ":" + PORT

	// Once a ServerConfig has been configured, connections can be accepted.
	listen, err := net.Listen(TYPE, address)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to listen on(%s): %s\n", address, err.Error())
		return nil, errors.New(errMsg)
	}
	fmt.Printf("Listening on %s...\n", PORT)

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			errmsg := fmt.Sprintf("error listening @%s error: %v", address, err.Error())
			return nil, errors.New(errmsg)
		}

		//sConn, chans, reqs, err := ssh.NewServerConn(conn, config)
		//if err != nil {
		//	errMsg := fmt.Sprintf("Error: %s", err.Error())
		//	return nil, errors.New(errMsg)
		//}
		//ssh.DiscardRequests(reqs)

		return conn, nil
    }
}
