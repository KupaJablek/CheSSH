package game

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func Host() {
    config := &ssh.ServerConfig {
        PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "foo" && string(pass) == "bar" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
    }

    key, _ := os.ReadFile("~/.ssh/id_rsa")
    signer, _ := ssh.ParsePrivateKey(key)
    config.AddHostKey(signer)
    
    address := "127.0.0.1:2200"

    listen, err := net.Listen("tcp", address)
    if err != nil {
        fmt.Println(err.Error()) 
        return
    }

    defer listen.Close()

    for {
        conn, err := listen.Accept()
        if err != nil {
            return
        }
        handleConn(conn)
    }
}

func handleConn(c net.Conn) {
   fmt.Println(c) 
}

