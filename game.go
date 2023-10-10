package main

import (
	"strconv"
	"strings"
)

type Game struct {
	board          [8][8]Piece
	game_over      bool
	current_player Player
	winner         Player
}

func (g *Game) move_piece(move string) (bool, string) {
	// not enough data in move
	if len(move) < 5 {
		return false, "Invalid format"
	}

	move_positions := strings.Split(move, "-")
	// not enough input || incorrect separator
	if len(move_positions) != 2 {
		return false, "Not enough data"
	}

	start_pos := decode_move(move_positions[0])
	if start_pos[0] == -1 || start_pos[1] == -1 {
		return false, "Invalid start position"
	}
	end_pos := decode_move(move_positions[1])
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
		if !g.validate_pawn(start_pos, end_pos) {
			return false, "pawn cannot move like that"
		}
		break
	case King:
		if !g.validate_king(start_pos, end_pos) {
			return false, "king cannot move like that"
		}
		break
	case Knight:
		if !g.validate_knight(start_pos, end_pos) {
			return false, "knight cannot move like that"
		}
		break
	case Rook:
		if !g.validate_orthogonal(start_pos, end_pos) {
			return false, "rook cannot move like that"
		}
		break
	case Bishop:
		if !g.validate_diagonal(start_pos, end_pos) {
			return false, "bishop cannot move like that"
		}
		break
	case Queen:
		if !g.validate_multi_direction(start_pos, end_pos) {
			return false, "queen cannot move like that"
		}
		break
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

	// check if pawn reached the other side of board... they get promoted to any piece except for pawn or king

	return true, ""
}

func (g *Game) promote_pawn(pawn [2]int) {

}

func (g *Game) end_turn() {
	if g.current_player == Player1 {
		g.current_player = Player2
	} else {
		g.current_player = Player1
	}
}

func decode_move(move string) [2]int {
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
		break
	case "b", "B":
		coordinates[1] = 1
		break
	case "c", "C":
		coordinates[1] = 2
		break
	case "d", "D":
		coordinates[1] = 3
		break
	case "e", "E":
		coordinates[1] = 4
		break
	case "f", "F":
		coordinates[1] = 5
		break
	case "g", "G":
		coordinates[1] = 6
		break
	case "h", "H":
		coordinates[1] = 7
		break
	default:
		break
	}

	return coordinates
}

func (g *Game) initialize_board() {
	g.board = [8][8]Piece{
		{
			{Player1, Rook, "R"}, {Player1, Knight, "H"}, {Player1, Bishop, "B"}, {Player1, King, "K"}, {Player1, Queen, "Q"}, {Player1, Bishop, "B"}, {Player1, Knight, "H"}, {Player1, Rook, "R"},
		},
		{
			{Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"}, {Player1, Pawn, "p"},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "}, {Empty, Blank, " "},
		},
		{
			{Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"}, {Player2, Pawn, "p"},
		},
		{
			{Player2, Rook, "R"}, {Player2, Knight, "H"}, {Player2, Bishop, "B"}, {Player2, King, "K"}, {Player2, Queen, "Q"}, {Player2, Bishop, "B"}, {Player2, Knight, "H"}, {Player2, Rook, "R"},
		},
	}
}
