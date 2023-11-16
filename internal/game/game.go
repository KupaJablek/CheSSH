package game

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/KupaJablek/CheSSH/internal/util"
)

type Game struct {
	board          [8][8]Piece
	game_over      bool
	current_player Player
	winner         Player

	conf util.Config
}

func MovePiece(g *Game, move string) (bool, string) {
	r, _ := regexp.Compile(`^[a-hA-H][1-8]-[a-hA-H][1-8]$`)

	//move = InputStringParser(move)

	if !r.MatchString(move) {
		return false, "Invalid move format: please use \"a#-a#\""
	}

	move_positions := strings.Split(move, "-")

	start_pos := DecodeMove(move_positions[0])
	if start_pos[0] == -1 || start_pos[1] == -1 {
		return false, "Invalid start position"
	}
	end_pos := DecodeMove(move_positions[1])
	if end_pos[0] == -1 || end_pos[1] == -1 {
		return false, "Invalid end position"
	}

	if start_pos == end_pos {
		return false, "No Position change"
	}

	start_piece := g.board[start_pos[0]][start_pos[1]]
	end_piece := g.board[end_pos[0]][end_pos[1]]

	if start_piece.player != g.current_player {
		return false, "Cannot move a piece thats not yours"
	}

	switch start_piece.piece_name {
	case Pawn:
		if !validate_pawn(g, start_pos, end_pos) {
			return false, "pawn cannot move like that"
		}
	case King:
		if !validate_king(g, start_pos, end_pos) {
			return false, "king cannot move like that"
		}
	case Knight:
		if !validate_knight(g, start_pos, end_pos) {
			return false, "knight cannot move like that"
		}
	case Rook:
		if !validate_orthogonal(g, start_pos, end_pos) {
			return false, "rook cannot move like that"
		}
	case Bishop:
		if !validate_diagonal(g, start_pos, end_pos) {
			return false, "bishop cannot move like that"
		}
	case Queen:
		if !validate_multi_direction(g, start_pos, end_pos) {
			return false, "queen cannot move like that"
		}
	default:
		return false, "Cannot move an empty space"
	}

	if start_piece.player == end_piece.player {
		return false, "cannot capture your own piece"
	}

	if end_piece.piece_name == King {
		g.game_over = true
		g.winner = g.current_player
	}

	g.board[end_pos[0]][end_pos[1]] = start_piece
	g.board[start_pos[0]][start_pos[1]] = Piece{Empty, Blank, " "}

	if start_piece.piece_name == Pawn {
		PromotePawn(g, end_pos)
	}

	return true, ""
}

func PromotePawn(g *Game, pawn [2]int) bool {
	if g.current_player == Player1 && pawn[0] != 7 {
		return false
	}
	if g.current_player == Player2 && pawn[0] != 0 {
		return false
	}

	fmt.Printf("Promote your pawn!\n")
	fmt.Printf("\tQueen: 'q'")
	fmt.Printf("\tBishop: 'b'")
	fmt.Printf("\tRook: 'r'")
	fmt.Printf("\tKnight: 'k'")
	fmt.Printf("Enter your choice:")

	validInput := false
	var promotedPiece Piece

	for !validInput {
		var userInput string
		fmt.Scanln(&userInput)

		switch userInput {
		case "q", "Q":
			promotedPiece = Piece{g.current_player, Queen, g.conf.Queen}
			validInput = true
		case "b", "B":
			promotedPiece = Piece{g.current_player, Bishop, g.conf.Bishop}
			validInput = true
		case "r", "R":
			promotedPiece = Piece{g.current_player, Rook, g.conf.Rook}
			validInput = true
		case "k", "K":
			promotedPiece = Piece{g.current_player, Knight, g.conf.Knight}
			validInput = true
		default:
			fmt.Printf("'%s' is not a valid choice, re-enter: ", userInput)
		}
	}
	g.board[pawn[0]][pawn[1]] = promotedPiece
	return true
}

func EndTurn(g *Game) {
	if g.current_player == Player1 {
		g.current_player = Player2
	} else {
		g.current_player = Player1
	}
}

func DecodeMove(move string) [2]int {
	coordinates := [2]int{-1, -1}
	row_pos, err := strconv.Atoi(string(move[1]))
	if err != nil || row_pos < 1 || row_pos > 8 {
		return coordinates
	}
	coordinates[0] = row_pos - 1

	// parse letter code
	switch string(move[0]) {
	case "a", "A":
		coordinates[1] = 0
	case "b", "B":
		coordinates[1] = 1
	case "c", "C":
		coordinates[1] = 2
	case "d", "D":
		coordinates[1] = 3
	case "e", "E":
		coordinates[1] = 4
	case "f", "F":
		coordinates[1] = 5
	case "g", "G":
		coordinates[1] = 6
	case "h", "H":
		coordinates[1] = 7
	default:
	}

	return coordinates
}

func InitializeBoard(g *Game) {
	g.board = [8][8]Piece{
		{
			{Player1, Rook, g.conf.Rook}, {Player1, Knight, g.conf.Knight}, {Player1, Bishop, g.conf.Bishop}, {Player1, King, g.conf.King},
			{Player1, Queen, g.conf.Queen}, {Player1, Bishop, g.conf.Bishop}, {Player1, Knight, g.conf.Knight}, {Player1, Rook, g.conf.Rook},
		},
		{
			{Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn},
			{Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn}, {Player1, Pawn, g.conf.Pawn},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn},
			{Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn}, {Player2, Pawn, g.conf.Pawn},
		},
		{
			{Player2, Rook, g.conf.Rook}, {Player2, Knight, g.conf.Knight}, {Player2, Bishop, g.conf.Bishop}, {Player2, King, g.conf.King},
			{Player2, Queen, g.conf.Queen}, {Player2, Bishop, g.conf.Bishop}, {Player2, Knight, g.conf.Knight}, {Player2, Rook, g.conf.Rook},
		},
	}
}
