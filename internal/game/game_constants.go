package game

type PieceType int

const (
	King PieceType = iota
	Queen
	Bishop
	Knight
	Rook
	Pawn
	Blank
)

type Player int

const (
	Empty Player = iota
	Player1
	Player2
)

type Piece struct {
	player     Player
	piece_name PieceType
	icon       string
}
