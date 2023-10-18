package main

import (
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		game.CreateHotseatGame()
		return
	}

	switch args[1] {
	case "--hotseat":
		game.CreateHotseatGame()
		break
	case "--host":
		game.HostSshLobby()
		break
	case "--join":
		game.JoinSshLobby()
		break
	case "--help":
	default:
		util.Help()
		break
	}
}
