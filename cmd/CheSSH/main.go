package main

import (
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		game.Create_Hotseat_Game()
		return
	}

	switch args[1] {
	case "--hotseat":
		game.Create_Hotseat_Game()
		break
	case "--host":
		game.Host_ssh_lobby()
		break
	case "--join":
		game.Join_ssh_lobby()
		break
	case "--help":
	default:
		util.Help()
		break
	}
}
