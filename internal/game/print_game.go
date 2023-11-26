package game

import (
	"fmt"

	"github.com/KupaJablek/CheSSH/internal/util"
	"github.com/fatih/color"
)

func PrintBoard(g *Game, c *util.Config, printToPlayer Player) {
	p1 := color.New(color.Attribute(g.conf.P1Colour), color.Bold)
	p2 := color.New(color.Attribute(g.conf.P2Colour), color.Bold)

	// board colour
	bc := color.New(color.Attribute(g.conf.BoardColour), color.Bold)

	var startCount int
	var limit int
	var increment int

	if g.current_player == Player1 {
		p1.Printf("Player %s's Turn\n", g.p1name)
	} else {
		p2.Printf("Player %s's Turn\n", g.p2name)
	}

	if printToPlayer == Player1 {
		startCount = 7
		limit = -1
		increment = -1
	} else {
		startCount = 0
		limit = 8
		increment = 1
	}

	bc.Println("\n   A B C D E F G H ")
	bc.Println("  +-+-+-+-+-+-+-+-+")
	for k := startCount; k != limit; k += increment {

		bc.Printf("%d |", k+1)
		for i := 0; i < 8; i++ {
			switch g.board[k][i].player {
			case Player1:
				p1.Printf("%s", g.board[k][i].icon)
				bc.Printf("|")
			case Player2:
				p2.Printf("%s", g.board[k][i].icon)
				bc.Printf("|")
			default:
				bc.Printf("%s|", g.board[k][i].icon)
			}
		}
		bc.Printf(" %d\n", k+1)
		bc.Println("  +-+-+-+-+-+-+-+-+")
	}
	bc.Println("   A B C D E F G H ")
	fmt.Println()
}
