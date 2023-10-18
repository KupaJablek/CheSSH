package game

import (
	"fmt"

	"github.com/KupaJablek/CheSSH/internal/util"
)

func CreateHotseatGame() {
	var g Game
	InitializeBoard(&g)
	g.current_player = Player1

	util.ClearTerminal()

	PrintBoard(&g)
	for !g.game_over {
		fmt.Println("enter chess coordinate ie: 'a1-a2' or n to surrender")
		var user_input string
		fmt.Scanln(&user_input)

		if user_input == "n" {
			if g.current_player == Player1 {
				g.winner = Player2
			} else {
				g.winner = Player1
			}
			g.game_over = true
			break
		}

		move_ok, err := MovePiece(&g, user_input)
		if move_ok {
			fmt.Println("ok")
		} else {
			fmt.Printf("Error: %s\n", err)
		}

		if move_ok {
			EndTurn(&g)
			util.ClearTerminal()
			PrintBoard(&g)
		}
	}

	util.ClearTerminal()
	fmt.Println("GAMEOVER")
	if g.winner == Player1 {
		fmt.Println("Player 1 is the Winner")
	} else {
		fmt.Println("Player 2 is the Winner")
	}
}

func HostSshLobby() {
	fmt.Println("NOT IMPLEMENTED YET")
}

func JoinSshLobby() {
	fmt.Println("NOT IMPLEMENTED YET")
}
