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

        if validMove(move, g) {
            fmt.Println("Invalid move")
            //fmt.Println(m)
            continue
        }
    }
}

func validMove(m string, g *Game) bool {
    if len(m) != 5 {
        return false
    }
    
    x1, x2 := int(m[0] - 'a'), int(m[3] - 'a')
    if x1 < 0 || x1 > 7 {
        return false
    } 
    if x2 < 0 || x2 > 7 {
        return false
    } 

    y1, err := strconv.Atoi(string(m[1]))
    y2, err2 := strconv.Atoi(string(m[4]))
    if err != nil || err2 != nil {
        return false
    }

    if y1 < 0 || y1 > 7 {
        return false
    } 
    if y2 < 0 || y2 > 7 {
        return false
    } 

    move := [][]int{{x1, y1},{x2, y2}}
    startPiece := g.board[x1][y1]

    switch startPiece.name {
    case 0: // blank
    return false
    case 1: // pawn
        if !validatePawn(move, g) {
            return false
        }
    case 2: // rook
        if !validateOrthogonal(move, g) {
            return false
        }
    case 3: // horse
    case 4: // bishop
    case 5: // king
    case 6: // queen
    }

    // update the board

    return true 
}

func validatePawn(m[][]int, g *Game) bool {
    yDiff := m[1][1] - m[0][1]
    if Abs(yDiff) != 1 && Abs(yDiff) != 2 {
        return false
    }

    xDiff := Abs(m[0][0] - m[1][0])
    if Abs(xDiff) != 1 && xDiff != 0 {
        return false
    }

    target := g.board[m[1][0]][m[1][1]]

    if yDiff == 2 {
        if target.player != 0 {
            return false
        }

        if g.p1Turn && m[0][1] != 1 {
            return false
        }

        if !g.p1Turn && m[0][1] != 6 {
            return false
        }
        return true
    }

    if xDiff == 0 && target.player != 0 {
        return false
    }

    if Abs(xDiff) == 1 {
        if g.p1Turn && target.player == 2 {
            return true
        }

        if !g.p1Turn && target.player == 1 {
            return true
        }
    }

    return false
}

func validateOrthogonal(m[][]int, g *Game) bool {
    target := g.board[m[1][0]][m[1][1]].player
    if g.p1Turn && target == 1 {
        return false
    }
    if !g.p1Turn && target == 2 {
        return false
    }

    yDiff := m[1][1] - m[0][1]
    xDiff := m[1][0] - m[0][0]

    if xDiff != 0 && yDiff != 0 {
        return false
    }
    if xDiff == 0 && yDiff == 0 {
        return false
    }

    if yDiff > 0 {
        for i := m[0][1]; i < m[1][1] - 1; i ++ {
            if g.board[i][m[0][0]].player != 0 {
                return false 
            }
        }
    }
    if yDiff < 0 {
        for i := m[0][1]; i > m[1][1] - 1; i -- {
            if g.board[i][m[0][0]].player != 0 {
                return false 
            }
        }
    }

    if xDiff > 0 {
        for i := m[0][0]; i < m[1][0] - 1; i ++ {
            if g.board[m[0][1]][i].player != 0 {
                return false 
            }
        }
    }
    if xDiff < 0 {
        for i := m[0][0]; i > m[1][0] - 1; i -- {
            if g.board[m[0][1]][i].player != 0 {
                return false 
            }
        }
    }

    return true
}

func PrintBoard(g *Game) {
    letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
    padding := -3

    fmt.Printf("%*s", padding - 1, "")
    for i := range letters {
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
