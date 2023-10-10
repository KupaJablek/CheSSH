package main

func (g *Game) validate_pawn(start [2]int, end [2]int) bool {
	y_diff := end[0] - start[0]
	pos_diff_x := absolute_val(end[1] - start[1])
	target := g.board[end[0]][end[1]]

	if g.current_player == Player1 {
		switch y_diff {
		case 1:
			break
		case 2:
			if start[0] != 1 { // index of second row from player1
				return false
			}
			break
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
			break
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

func (g *Game) validate_orthogonal(start [2]int, end [2]int) bool {

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

func (g *Game) validate_diagonal(start [2]int, end [2]int) bool {

	if start[0] == end[0] || start[1] == end[1] {
		return false
	}

	x_diff := end[1] - start[1]
	y_diff := end[0] - start[0]

	if absolute_val(y_diff) != absolute_val(x_diff) {
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

func (g *Game) validate_multi_direction(start [2]int, end [2]int) bool {
	if !g.validate_diagonal(start, end) && !g.validate_orthogonal(start, end) {
		return false
	}
	return true
}

func (g *Game) validate_knight(start [2]int, end [2]int) bool {
	ydiff := absolute_val(end[0] - start[0])
	xdiff := absolute_val(end[1] - start[1])
	if ydiff == 2 && xdiff == 1 {
		return true
	} else if ydiff == 1 && xdiff == 2 {
		return true
	}
	return false
}

func (g *Game) validate_king(start [2]int, end [2]int) bool {
	if absolute_val(end[0]-start[0]) <= 1 && absolute_val(end[1]-start[1]) <= 1 {
		return g.validate_multi_direction(start, end)
	}
	return false
}
