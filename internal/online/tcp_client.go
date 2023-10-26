package online

import (
	"errors"
	"fmt"
	"net"
)

func JoinTCP(HOST string, PORT string, TYPE string) (*net.TCPConn, error) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		errorMessage := fmt.Sprintf("Cannot connect to: %s", HOST+":"+PORT)
		fmt.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		errorMessage := fmt.Sprintf("Dial failed: %s", err.Error())
		fmt.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	return conn, nil
}
