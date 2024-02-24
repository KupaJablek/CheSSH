package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	pk, _ := os.ReadFile("") //PATH to private key
	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		panic(err)
	}

	hostkeyCallback, err := knownhosts.New("") //PATH to .knownhosts
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

	client, err := ssh.Dial("tcp", "", config) //IP to connect

	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	//Example command of what can be run. Should look into server config with no authorization asw 
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	defer session.Close()
}
