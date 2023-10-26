package main

import (
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

// default values for host ip and port #
const (
	HOST = "0.0.0.0"
	PORT = "2200"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		game.CreateHotseatGame()
		return
	}

	switch args[1] {
	case "--hotseat": // --hotseat -p1 player1name -p2 player2name
		game.CreateHotseatGame()

	case "--host": // --host -ip 0.0.0.0 -p 2200 -u sessionID
		game.HostLobby(HOST, PORT)

	case "--join": // -- join -ip 0.0.0.0 -p 2200 -u sessionID
		game.JoinLobby(HOST, PORT)

	case "--help":
		util.Help()

	default:
		util.Help()
	}
}
