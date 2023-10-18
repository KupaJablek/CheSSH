package game

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBoard(g *Game) {
	blue := color.New(color.FgBlue, color.Bold)
	red := color.New(color.FgRed, color.Bold)

	var startCount int
	var limit int
	var increment int

	if g.current_player == Player1 {
		startCount = 7
		limit = -1
		increment = -1
		blue.Printf("Player 1's Turn\n")
	} else {
		startCount = 0
		limit = 8
		increment = 1
		red.Printf("Player 2's Turn\n")
	}

	fmt.Println("\n   A B C D E F G H ")
	fmt.Println("  +-+-+-+-+-+-+-+-+")
	for k := startCount; k != limit; k += increment {

		fmt.Printf("%d |", k+1)
		for i := 0; i < 8; i++ {
			switch g.board[k][i].player {
			case Player1:
				blue.Printf("%s", g.board[k][i].icon)
				fmt.Printf("|")
				break
			case Player2:
				red.Printf("%s", g.board[k][i].icon)
				fmt.Printf("|")
				break
			default:
				fmt.Printf("%s|", g.board[k][i].icon)
				break
			}
		}
		fmt.Printf(" %d\n", k+1)
		fmt.Println("  +-+-+-+-+-+-+-+-+")
	}
	fmt.Println("   A B C D E F G H ")
	fmt.Println()
}
