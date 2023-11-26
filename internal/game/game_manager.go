package game

import (
	"bytes"
	"fmt"

	"github.com/KupaJablek/CheSSH/internal/online"
	"github.com/KupaJablek/CheSSH/internal/util"
)

func CreateHotseatGame(p1name, p2name string) {
	var g Game
	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)
	g.conf = conf

	g.p1name = p1name
	g.p2name = p2name

	InitializeBoard(&g)
	g.current_player = Player1

	util.ClearTerminal()
	PrintBoard(&g, &conf, g.current_player)

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
		PrintBoard(&g, &conf, g.current_player)
	}
	ShowGameOverScreen(&g)
}

func HostLobby(HOST, PORT, USER, PASSWORD string) {
	conn, err := online.HostTCP(HOST, PORT, "tcp")
	//conn, err := online.CreateSSHServer(HOST, PORT, "tcp")
	if err != nil {
		fmt.Println("Error hosting server: ", err.Error())
		return
	}

	// clear messages from before connection
	util.ClearTerminal()

	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)

	var g Game
	g.p1name = "Player 1"
	g.p2name = "Player 2"

	g.conf = conf
	InitializeBoard(&g)
	g.current_player = Player1

	for {
		// server player's turn

		// send data to client
		PrintBoard(&g, &conf, Player1)
		move := GetPlayerMove(&g)
		fmt.Fprint(conn, move)
		EndTurn(&g)
		if g.game_over {
			conn.Close()
			break
		}

		// client players turn
		util.ClearTerminal()
		PrintBoard(&g, &conf, Player1)

		// recieve data from client
		buffer := make([]byte, 5)
		_, err := conn.Read(buffer)
		cleanData := bytes.Trim(buffer, "\x00")
		var parsedMove = string(cleanData)
		if err != nil {
			fmt.Println("Error reading data from server: ", err.Error())
		}

		if parsedMove == "n" {
			// user has surrendered
			g.game_over = true
			g.winner = Player1
			conn.Close()
			break
		}

		// sync client side move with local game
		MovePiece(&g, parsedMove)
		EndTurn(&g)
		if g.game_over {
			conn.Close()
			break
		}
	}
	ShowGameOverScreen(&g)
}

func JoinLobby(HOST, PORT, USER string) {
	conn, err := online.JoinTCP(HOST, PORT, "tcp")
	//conn, err := online.JoinSSHLobby(HOST, PORT, "tcp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conf, _ := util.LoadConfig()
	util.InitDefault(&conf)

	var g Game
	g.p1name = "Player 1"
	g.p2name = "Player 2"
	g.conf = conf
	InitializeBoard(&g)
	g.current_player = Player1

	// while connection is open
	for {
		// server player's turn
		util.ClearTerminal()

		PrintBoard(&g, &conf, Player2)

		// recieve data from server
		buffer := make([]byte, 5)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data from server: ", err.Error())
		}

		cleanData := bytes.Trim(buffer, "\x00")
		var parsedMove = string(cleanData)

		// sync server side move with local game
		if parsedMove == "n" {
			// user has surrendered
			g.game_over = true
			g.winner = Player2
			conn.Close()
			break
		}
		MovePiece(&g, parsedMove)
		EndTurn(&g)

		if g.game_over {
			conn.Close()
			break
		}

		util.ClearTerminal()

		// send data to server
		PrintBoard(&g, &conf, Player2)
		move := GetPlayerMove(&g)
		EndTurn(&g)
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
	fmt.Print("!! GAMEOVER !!\n\n")
	if g.winner == Player1 {
		fmt.Printf("%s is the Winner\n", g.p1name)
	} else {
		fmt.Printf("%s is the Winner\n", g.p2name)
	}
}

func GetPlayerMove(g *Game) string {
	var userInput string
	var validMove = false

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
