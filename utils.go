package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func absolute_val(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func help() {
	fmt.Println("CheSSH is a tool for playing local or online chess using SSH")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\tchessh <command> [arguments]")
	fmt.Println("")
	fmt.Println("The Commands are:")
	fmt.Println("\t--hotseat")
	fmt.Println("\t--host")
	fmt.Println("\t--join")
}

func clear_terminal() {
	platform := runtime.GOOS
	if platform == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func print_board(g Game) {
	blue := color.New(color.FgBlue, color.Bold)
	red := color.New(color.FgRed, color.Bold)

	var startCount int
	var limit int
	var increment int

	if g.current_player == Player1 {
		startCount = 7
		limit = -1
		increment = -1
		fmt.Println("Player 1's Turn")
	} else {
		startCount = 0
		limit = 8
		increment = 1
		fmt.Println("Player 2's Turn")
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
