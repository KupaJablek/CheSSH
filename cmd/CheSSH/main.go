package main

import (
	"fmt"
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

// default values for roomID, host ip, port #, etc...
var HOST = "127.0.0.1"
var PORT = "2200"
var PASSWORD = ""
var USERID = "Chess"
var P1 = "Player 1"
var P2 = "Player 2"

func main() {
	args := os.Args

	if len(args) == 1 || len(args)%2 != 0 {
		util.Help()
		return
	}

	if len(args) > 2 && !InputStringParser(args) {
		util.Help()
		return
	}

	switch args[1] {
	case "--hotseat": // --hotseat -p1 player1name -p2 player2name
		game.CreateHotseatGame(P1, P2)

	case "--host": // --host -ip 0.0.0.0 -p 2200 -u userID
		game.HostLobby(HOST, PORT, USERID, PASSWORD)

	case "--join": // -- join -ip 0.0.0.0 -p 2200 -u userID
		game.JoinLobby(HOST, PORT, USERID)

	case "--help":
		util.Help()

	default:
		util.Help()
	}
}

func InputStringParser(input []string) bool {
	for i := 2; i < len(input)-1; i++ {
		// check if last element in args is a command flag with no param
		if input[i][0] != '-' && i == len(input)-1 {
			fmt.Printf("'%s' is missing a value\n", input[i])
			return false
		}

		// check if flag has a valid param
		if input[i+1][0] == '-' {
			fmt.Printf("'%s' is missing a value\n", input[i])
			return false
		}

		if input[1] == "--hotseat" {
			switch input[i] {

			case "-p1":
				i++
				P1 = input[i]
			case "-p2":
				i++
				P2 = input[i]
			default:
				fmt.Printf("%s is not a valid command flag for %s\n", input[i], input[1])
				return false
			}
		} else {
			switch input[i] {
			case "-ip":
				i++
				HOST = input[i]
			case "-p", "-port":
				i++
				PORT = input[i]
			case "-u":
				i++
				USERID = input[i]
			default:
				fmt.Printf("%s is not a valid command flag for %s\n", input[i], input[1])
				return false
			}
		}
	}
	return true
}
