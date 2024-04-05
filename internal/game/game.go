package game

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type Piece struct {
    name, player int
}

type Game struct {
    board [][]Piece
    p1Turn, gameOver, p1Winner bool
}

/*
0 : blank
1 : pawn
2 : rook
3 : horse
4 : bishop
5 : king
6 : queen
*/

func InitializeBoard(g *Game) {
    board := make([][]Piece, 8)

    for i := 2; i < 6; i++ {
        row := make([]Piece, 8)
        for r := range row {
            row[r] = Piece{0,0}
        }
        board[i] = row
    }

    p1, p2 := make([]Piece, 8), make([]Piece, 8)
    for r := range p1 {
        p1[r] = Piece{1, 1}
        p2[r] = Piece{1, 2}
    }
    board[1], board[6] = p1, p2

    f1 := []Piece{{2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {4, 1}, {3, 1}, {2, 1}}
    f2 := []Piece{{2, 2}, {3, 2}, {4, 2}, {5, 2}, {6, 2}, {4, 2}, {3, 2}, {2, 2}}
    board[0], board[7] = f1, f2
    g.board = board
}

func TakeTurn(g *Game) {
    var move string
    var valid bool
    for !valid {
        fmt.Println("Enter your move or h for help:")
        fmt.Scanln(&move)

        if move == "resign" || move == "ff" {
            g.gameOver = true
            g.p1Winner = g.p1Turn == false
            return
        }

        if m := validMove(move); m == nil {
            fmt.Println("Invalid move")
            fmt.Println(m)
            continue
        }
    }
}

func validMove(m string) [][]int {
    if len(m) != 5 {
        return nil
    }
    
    x1 := int(m[0] - 'a')
    x2 := int(m[3] - 'a')
    if x1 < 0 || x1 > 7 {
        return nil
    } 
    if x2 < 0 || x2 > 7 {
        return nil
    } 

    y1, err := strconv.Atoi(string(m[1]))
    if err != nil {
        return nil
    }
    y2, err := strconv.Atoi(string(m[4]))
    if err != nil {
        return nil
    }

    // check for piece specific validation

    return [][]int{{x1, y1},{x2, y2}}
}

func PrintBoard(g *Game) {
    letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
    padding := -3

    fmt.Printf("%*s", padding - 1, "")
    for i := 0; i < 8; i++ {
        fmt.Printf("%*c", padding - 1, letters[i])
    }
    fmt.Printf("\n%*s", padding - 1, "")
    fmt.Println("- - - - - - - - - - - - - - -")
    
    var si, lim, inc int
    if !g.p1Turn {
        lim = 8
        inc = 1
    } else {
        si = 7
        lim = -1
        inc = -1
    }

	p1 := color.New(color.Attribute(34), color.Bold)
	p2 := color.New(color.Attribute(31), color.Bold)

    for i := si; i != lim; i += inc {
        fmt.Printf("%d", i + 1)
        fmt.Printf("%*s", padding, "")
        for k := 0; k < 8; k++ {
            p := g.board[i][k] 
            piece := ""
            switch p.name {
            case 0:
                piece = " "
            case 1:
                piece = "p"
            case 2:
                piece = "r"
            case 3:
                piece = "k"
            case 4:
                piece = "b"
            case 5:
                piece = "K"
            case 6:
                piece = "Q"
            } 
            if p.player == 1 {
                p1.Printf("%s", piece)
            } else if p.player == 2 {
                p2.Printf("%s", piece)
            } else {
                fmt.Printf(piece)
            }
            fmt.Printf("%*s", -padding, "")
        }
        fmt.Printf("%d\n", i + 1)
    }

    fmt.Printf("%*s", padding - 1, "")
    fmt.Println("- - - - - - - - - - - - - - -")
    fmt.Printf("%*s", padding - 1, "")
    for i := 0; i < 8; i++ {
        fmt.Printf("%*c", padding - 1, letters[i])
    }
    fmt.Println()
}
