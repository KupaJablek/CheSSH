package main

import "fmt"

func create_hotseat_game() {
	var g Game
	g.initialize_board()
	g.current_player = Player1

	clear_terminal()

	print_board(g)
	for !g.game_over {
		fmt.Println("enter chess coordinate ie: 'a1-a2' or n to end game")
		var user_input string
		fmt.Scanln(&user_input)

		if user_input == "n" {
			return
		}

		move_ok, err := g.move_piece(user_input)
		if move_ok {
			fmt.Println("ok")
		} else {
			fmt.Printf("Error: %s\n", err)
		}

		if move_ok {
			g.end_turn()
			clear_terminal()
			print_board(g)
		}
	}
	fmt.Println("GAMEOVER")
	if g.winner == Player1 {
		fmt.Println("Player 1 is the Winner")
	} else {
		fmt.Println("Player 2 is the Winner")
	}
}

func host_ssh_lobby() {
	fmt.Println("NOT IMPLEMENTED YET")
}

func join_ssh_lobby() {
	fmt.Println("NOT IMPLEMENTED YET")
}
