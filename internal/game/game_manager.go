package game

import (
	"fmt"

	"github.com/KupaJablek/CheSSH/internal/online"
	"github.com/KupaJablek/CheSSH/internal/util"
)

func CreateHotseatGame(p1name, p2name string) {
	var g Game
	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)
	g.conf = conf

	InitializeBoard(&g)
	g.current_player = Player1

	util.ClearTerminal()
	PrintBoard(&g, &conf)

	for !g.game_over {
		move := GetPlayerMove(&g)

		EndTurn(&g)
		util.ClearTerminal()
		if g.current_player == Player1 {
			fmt.Printf("%s moved: %v", p1name, move)
		} else {
			fmt.Printf("%s moved: %v", p2name, move)
		}
		fmt.Println("")
		PrintBoard(&g, &conf)
	}
	ShowGameOverScreen(&g)
}

func HostLobby(HOST, PORT, USER, PASSWORD string) {
	fmt.Print("NOT FULLY IMPLEMENTED YET\n\n")
	conn, err := online.HostTCP(HOST, PORT, "tcp")
	//conn, err := online.CreateSSHServer(HOST, PORT, "tcp")
	if err != nil {
		fmt.Println("Error hosting server: ", err.Error())
		return
	}

	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)

	var g Game
	g.conf = conf
	InitializeBoard(&g)
	g.current_player = Player1

	for {
		// server player's turn
		fmt.Printf("It's your turn!\n")

		// send data to client
		PrintBoard(&g, &conf)
		move := GetPlayerMove(&g)
		fmt.Fprint(conn, move)
		if g.game_over {
			conn.Close()
			break
		}

		// client players turn
		util.ClearTerminal()
		fmt.Printf("It is Player 2's turn")

		// recieve data from client
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data from server: ", err.Error())
		}

		// sync client side move with local game
		MovePiece(&g, string(buffer))
		EndTurn(&g)
		if g.game_over {
			conn.Close()
			break
		}
	}
	ShowGameOverScreen(&g)
}

func JoinLobby(HOST, PORT, USER string) {
	fmt.Print("NOT FULLY IMPLEMENTED YET\n\n")
	conn, err := online.JoinTCP(HOST, PORT, "tcp")
	//conn, err := online.JoinSSHLobby(HOST, PORT, "tcp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)

	var g Game
	g.conf = conf
	InitializeBoard(&g)
	g.current_player = Player1

	// while connection is open
	for {
		// server player's turn
		util.ClearTerminal()
		fmt.Printf("It is Player 1's turn\n")

		// recieve data from server
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data from server: ", err.Error())
		}

		// sync server side move with local game
		MovePiece(&g, string(buffer))
		EndTurn(&g)

		if g.game_over {
			conn.Close()
			break
		}

		util.ClearTerminal()
		fmt.Printf("It's your turn!\n")

		// send data to server
		PrintBoard(&g, &conf)
		move := GetPlayerMove(&g)
		fmt.Fprint(conn, move)
		if g.game_over {
			conn.Close()
			break
		}
	}

	ShowGameOverScreen(&g)
}

func ShowGameOverScreen(g *Game) {
	util.ClearTerminal()
	fmt.Println("GAMEOVER")
	if g.winner == Player1 {
		fmt.Println("Player 1 is the Winner")
	} else {
		fmt.Println("Player 2 is the Winner")
	}
}

func GetPlayerMove(g *Game) string {
	var userInput string
	validMove := false

	for !validMove {
		fmt.Println("enter chess coordinate ie: 'a1-a2' or n to surrender")
		fmt.Scanln(&userInput)

		if userInput == "n" {
			if g.current_player == Player1 {
				g.winner = Player2
			} else {
				g.winner = Player1
			}
			g.game_over = true
			break
		}

		err := ""
		validMove, err = MovePiece(g, userInput)
		if err != "" {
			fmt.Printf("Error: %s\n", err)
		}
	}

	return userInput
}
