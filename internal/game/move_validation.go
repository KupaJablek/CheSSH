package game

import "github.com/KupaJablek/CheSSH/internal/util"

func validate_pawn(g *Game, start [2]int, end [2]int) bool {
	y_diff := end[0] - start[0]
	pos_diff_x := util.AbsoluteVal(end[1] - start[1])
	target := g.board[end[0]][end[1]]

	if g.current_player == Player1 {
		switch y_diff {
		case 1:
			break
		case 2:
			if start[0] != 1 { // index of second row from player1
				return false
			}
		default:
			return false
		}
	}
	if g.current_player == Player2 { // index of second row from player2
		switch y_diff {
		case -1:
			break
		case -2:
			if start[0] != 6 {
				return false
			}
		default:
			return false
		}
	}

	if pos_diff_x == 0 && target.player != Empty {
		return false
	} else if pos_diff_x == 1 {
		if g.current_player == Player1 && g.board[end[0]][end[1]].player != Player2 {
			return false
		} else if g.current_player == Player2 && g.board[end[0]][end[1]].player != Player1 {
			return false
		}
	}

	return true
}

func validate_orthogonal(g *Game, start [2]int, end [2]int) bool {

	y_diff := end[0] - start[0]
	x_diff := end[1] - start[1]
	if start[0] != end[0] && start[1] != end[1] {
		return false
	}

	if y_diff > 0 {
		for i := start[0] + 1; i < end[0]; i++ {
			if g.board[i][start[1]].player != Empty && i != end[0] {
				return false
			}
		}
	}

	if y_diff < 0 {
		for i := start[0] - 1; i > end[0]; i-- {
			if g.board[i][start[1]].player != Empty && i != end[0] {
				return false
			}
		}
	}

	if x_diff > 0 {
		for i := start[1] + 1; i < end[1]; i++ {
			if g.board[start[0]][i].player != Empty && i != end[1] {
				return false
			}
		}
	}

	if x_diff < 0 {
		for i := start[1] - 1; i > end[1]; i-- {
			if g.board[start[0]][i].player != Empty && i != end[1] {
				return false
			}
		}
	}

	return true
}

func validate_diagonal(g *Game, start [2]int, end [2]int) bool {

	if start[0] == end[0] || start[1] == end[1] {
		return false
	}

	x_diff := end[1] - start[1]
	y_diff := end[0] - start[0]

	if util.AbsoluteVal(y_diff) != util.AbsoluteVal(x_diff) {
		return false
	}

	if y_diff < 0 && x_diff < 0 {
		for i := -1; i >= y_diff; i-- {
			if g.board[start[0]+i][start[1]+i].player != Empty && start[0]-i != end[0] {
				return false
			}
		}
	} else if y_diff < 0 && x_diff > 0 {
		for i := 1; i <= x_diff; i++ {
			if g.board[start[0]-i][start[1]+i].player != Empty && i != x_diff {
				return false
			}
		}

	} else if y_diff > 0 && x_diff < 0 {
		for i := 1; i <= y_diff; i++ {
			if g.board[start[0]+i][start[1]-i].player != Empty && i != y_diff {
				return false
			}
		}
	} else if y_diff > 0 && x_diff > 0 {
		for i := 1; i <= y_diff; i++ {
			if g.board[start[0]+i][start[1]+i].player != Empty && i != y_diff {
				return false
			}
		}
	}
	return true
}

func validate_multi_direction(g *Game, start [2]int, end [2]int) bool {
	if !validate_diagonal(g, start, end) && !validate_orthogonal(g, start, end) {
		return false
	}
	return true
}

func validate_knight(g *Game, start [2]int, end [2]int) bool {
	ydiff := util.AbsoluteVal(end[0] - start[0])
	xdiff := util.AbsoluteVal(end[1] - start[1])
	if ydiff == 2 && xdiff == 1 {
		return true
	} else if ydiff == 1 && xdiff == 2 {
		return true
	}
	return false
}

func validate_king(g *Game, start [2]int, end [2]int) bool {
	if util.AbsoluteVal(end[0]-start[0]) <= 1 && util.AbsoluteVal(end[1]-start[1]) <= 1 {
		return validate_multi_direction(g, start, end)
	}
	return false
}
