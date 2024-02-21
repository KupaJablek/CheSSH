package main

import (
	"fmt"
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
	"github.com/KupaJablek/CheSSH/internal/util"
)

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

	if !InputStringParser(args) {
		util.Help()
		return
	}

	switch args[1] {
	case "--hotseat":
		game.CreateHotseatGame(P1, P2)

	case "--host":
		game.HostLobby(HOST, PORT, USERID, PASSWORD)

	case "--join":
		game.JoinLobby(HOST, PORT, USERID)

	default:
		util.Help()
	}
}

func InputStringParser(input []string) bool {
	for i := 2; i < len(input)-1; i++ {
		if input[i][0] != '-' && i == len(input)-1 {
			fmt.Printf("'%s' is missing a value\n", input[i])
			return false
		}

		if input[i+1][0] == '-' {
			fmt.Printf("'%s' is missing a value\n", input[i])
			return false
		}

        i++
		if input[1] == "--hotseat" {
			switch input[i] {
			case "-p1":
				P1 = input[i]
			case "-p2":
				P2 = input[i]
			default:
				fmt.Printf("%s is not a valid command flag for %s\n", input[i], input[1])
				return false
			}
		} else {
			switch input[i] {
			case "-ip":
				HOST = input[i]
			case "-p", "-port":
				PORT = input[i]
			case "-u":
				USERID = input[i]
			default:
				fmt.Printf("%s is not a valid command flag for %s\n", input[i], input[1])
				return false
			}
		}
	}
	return true
}
