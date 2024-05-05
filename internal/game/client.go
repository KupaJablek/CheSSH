package game

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func Join() {
    pk, _ := os.ReadFile(currentUserKeyPath()) //PATH to private key
    signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		panic(err)
	}

	hostkeyCallback, err := knownhosts.New(knownHostsPath()) //PATH to .knownhosts
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: "", //Username to connect
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostkeyCallback,
	}

    address := "127.0.0.1" + ":" + "2200"
	client, err := ssh.Dial("tcp", address, config) //IP to connect

	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	} else {
        fmt.Println("session created")
    }

	defer session.Close()

    for {
        uInput := ""
        fmt.Scanln(&uInput)

        b := []byte(uInput)
        
        session.Stdout.Write(b)
    }
}
