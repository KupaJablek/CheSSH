package main

import (
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

// default values for host ip and port #

var HOST = "0.0.0.0"
var PORT = "2200"
var sessionID = "sessionID"
var p1 = "Player 1"
var p2 = "Player 2"

func main() {
	args := os.Args
	if len(args) < 2 {
		game.CreateHotseatGame()
		return
	}

	if len(args) > 2 {
		if InputStringParser(args) == "fail" {
			//bad input
		}
	}

	switch args[1] {
	case "--hotseat": // --hotseat -p1 player1name -p2 player2name
		game.CreateHotseatGame()
		//add createhotsetgame with p1 & p2 params for name
		//game.CreateHotSeatGame(p1, p2)

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

func InputStringParser(input []string) string {
	result := "success"
	for i := 2; i < len(input); i++ {
		switch input[i] {
		case "-ip":
			i++
			HOST = input[i]
		case "-p", "-port":
			i++
			PORT = input[i]
		case "-u":
			i++
			sessionID = input[i]
		case "-p1":
			i++
			p1 = input[i]
		case "-p2":
			i++
			p2 = input[i]
		default:
			return "fail"
		}
	}

	return result
}
