package online

import (
	"errors"
	"fmt"
	"net"
)

func HostTCP(HOST string, PORT string, TYPE string) (net.Conn, error) {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	fmt.Printf("Lobby is open on %s", (HOST + ":" + PORT))

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to listen on(%s): %s\n", HOST+":"+PORT, err.Error())
		fmt.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return nil, errors.New("Error: " + err.Error())
		}
		return conn, nil
	}
}
