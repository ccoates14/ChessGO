package board

import "ChessGo/player"
import "math"

func ValidQueenMove(move *player.Move, gameBoard *Board) bool {
	bishopMovement := move.FromCol != move.ToCol && move.FromRow != move.ToRow 

	if bishopMovement {
		return ValidBishopMove(move, gameBoard)
	} 

	return ValidRookMove(move, gameBoard)
}

func ValidPawnMove(move *player.Move, gameBoard *Board) bool {
	//on first move can move forward two 
	//else can move forward one
	// can attack diagonal left or right one space if there is something to attack
	var toPiece = getBoardString(move.ToCol, move.ToRow, gameBoard)

	if toPiece != Empty {
		return false
	}

	//is moving forward
	if move.FromCol == move.ToCol {
		var difference = math.Abs(float64(move.ToRow - move.FromRow))

		if difference == 2 { //check if first move
			if move.WhitePlayer {
				if move.FromRow == 6 { //0 based
					//check piece would be jumped
					var skippedPosition = getBoardString(move.FromCol, move.ToRow + 1, gameBoard)

					if skippedPosition != Empty {
						return false
					} else {
						return true
					}
				}  else {
					return false
				}
			} else {
				//I could refactor this later...but duplication isn't always a bad thing
				if move.FromRow == 1 { //0 based
					//check piece would be jumped
					var skippedPosition = getBoardString(move.FromCol, move.ToRow - 1, gameBoard)

					if skippedPosition != Empty {
						return false
					} else {
						return true
					}
				}  else {
					return false
				}
			}
		} else if difference != 1 {
			return false
		}
	} else { //else is attacking
		//it can only move up and to the left or right one

		var differenceCol = math.Abs(float64(move.ToCol - move.FromCol))
		var differenceRow = math.Abs(float64(move.ToRow - move.FromRow))

		if differenceCol != 1 || differenceRow != 1 {
			return false
		}
	}

	return true
}

func ValidKingMove(move *player.Move, gameBoard *Board) bool {
	verticalDistance := math.Abs(float64(move.FromRow - move.ToRow))
	horizontalDistance := math.Abs(float64(move.FromCol - move.ToCol))

	if horizontalDistance <= 1 && verticalDistance <= 1 {
		return true
	}

	return false //if here not a valid distance
}

func ValidRookMove(move *player.Move, gameBoard *Board) bool {

	if move.FromCol == move.ToCol {
		//can move up and down
		//basically iterate through till we hit destination and check that we don't hit a piece along the way except for at the end
		var direction = 1 //dowwn

		if move.FromRow > move.ToRow { // up
			direction = -1
		}
	
		var distance = math.Abs(float64(move.FromRow - move.ToRow))
	
		for i := 0; i < int(distance); i++ {
			//check if current position contains a piece
			if getBoardString(move.FromCol, move.FromRow + (direction * i), gameBoard) != Empty {
				return false
			}
		}
	} else if move.FromRow == move.ToRow { //left right
		var direction = 1 //right

		if move.FromCol > move.ToCol { // left
			direction = -1
		}
	
		var distance = math.Abs(float64(move.FromCol - move.ToCol))
	
		for i := 0; i < int(distance); i++ {
			//check if current position contains a piece
			if getBoardString(move.FromCol + (direction * i), move.FromRow, gameBoard) != Empty {
				return false
			}
		}
	}

	return false //if here then they aren't move up down left right
}

func ValidKnightMove(move *player.Move, gameBoard *Board) bool {
	verticalDistance := math.Abs(float64(move.FromRow - move.ToRow))
	horizontalDistance := math.Abs(float64(move.FromCol - move.ToCol))

	if horizontalDistance == 1 && verticalDistance == 2 {
		return true
	}

	return false //if here not a valid L shape
}

func ValidBishopMove(move *player.Move, gameBoard *Board) bool {
	//move both vertical and horizontal at same time
	if move.FromCol == move.ToCol || move.FromRow == move.ToRow {
		return false
	}

	//then check that the path it is moving on does not hit a piece prior to end destination

	distance := math.Abs(float64(move.FromRow - move.ToRow))

	h_direction := 1 //right
	v_direction := 1 //down

	if move.FromCol > move.ToCol {
		h_direction = -1
	} 

	if move.FromRow > move.ToRow {
		v_direction = -1
	}

	for i := 0; i < int(distance); i++ {
		if getBoardString(move.FromCol + (h_direction * i), move.FromRow + (v_direction * i), gameBoard) != Empty {
			return false
		}
	}

	return true
}

func PawnInCheck(move *player.Move, gameBoard *Board) bool {
	//at this point with all these potential in check methods we assume
	//the move has already been performed
	//use the to xy

	verticalPawnDirection := -1

	if !move.WhitePlayer {
		verticalPawnDirection = 1
	}

	if move.ToRow + verticalPawnDirection >= 0 && move.ToRow + verticalPawnDirection < HEIGHT {
		//check left
		if move.ToCol - 1 >= 0 {
			var potentialPawn = getBoardString(move.ToCol - 1, move.ToRow + verticalPawnDirection, gameBoard)
			if (move.WhitePlayer && potentialPawn == BPawn) || (!move.WhitePlayer && potentialPawn == WPawn) {
				return true
			}
		}

		//check right
		if move.ToCol + 1 < WIDTH {
			var potentialPawn = getBoardString(move.ToCol + 1, move.ToRow + verticalPawnDirection, gameBoard)
			if (move.WhitePlayer && potentialPawn == BPawn) || (!move.WhitePlayer && potentialPawn == WPawn) {
				return true
			}
		}
	}

	return false
}

func RookInCheck(move *player.Move, gameBoard *Board) bool {
	//up down left right

	//starting from king position
	//move out in a cross shape and see if you encounter a rook of the opposing team

	var up, down, left, right = move.ToRow - 1, move.ToRow + 1, move.ToCol - 1, move.ToCol + 1

	for up - move.ToRow >= 0 {
		var potentialRook = getBoardString(move.ToCol, up - move.ToRow, gameBoard)

		if (potentialRook == BRook && move.WhitePlayer) || (potentialRook == WRook && !move.WhitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}
	}

	for down + move.ToRow < HEIGHT {
		var potentialRook = getBoardString(move.ToCol, down + move.ToRow, gameBoard)

		if (potentialRook == BRook && move.WhitePlayer) || (potentialRook == WRook && !move.WhitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}
	}

	for left - move.ToCol >= 0 {
		var potentialRook = getBoardString(left - move.ToCol, move.ToRow, gameBoard)

		if (potentialRook == BRook && move.WhitePlayer) || (potentialRook == WRook && !move.WhitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}
	}

	for right + move.ToCol < WIDTH {
		var potentialRook = getBoardString(right + move.ToCol, move.ToRow, gameBoard)

		if (potentialRook == BRook && move.WhitePlayer) || (potentialRook == WRook && !move.WhitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}
	}

	return false
}

func BishopInCheck(move *player.Move, gameBoard *Board) bool {
	col, row := move.ToCol + 1, move.ToRow - 1

	//right down - zero based
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && move.WhitePlayer) || (potentialBishop == WBishop && !move.WhitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row--
	}

	col, row = move.ToCol - 1, move.ToRow - 1
	//left down --zero based
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && move.WhitePlayer) || (potentialBishop == WBishop && !move.WhitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row--
	}



	//diagonal right left up - zero based
	col, row = move.ToCol + 1, move.ToRow + 1
	//right up
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && move.WhitePlayer) || (potentialBishop == WBishop && !move.WhitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row++
	}

	col, row = move.ToCol - 1, move.ToRow + 1
	//left up zero based
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && move.WhitePlayer) || (potentialBishop == WBishop && !move.WhitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row++
	}



	return false
}

func KingInCheck(move *player.Move, gameBoard *Board) bool {

	if move.ToCol - 1 >= 0 {
		potentialKing := getBoardString(move.ToCol - 1, move.ToRow, gameBoard)

		if (potentialKing == BKing && move.WhitePlayer) || (potentialKing == WKing && !move.WhitePlayer) {
			return true
		}
	}

	if move.ToCol + 1 < WIDTH {
		potentialKing := getBoardString(move.ToCol + 1, move.ToRow, gameBoard)

		if (potentialKing == BKing && move.WhitePlayer) || (potentialKing == WKing && !move.WhitePlayer) {
			return true
		}
	}

	if move.ToRow - 1 >= 0 {
		potentialKing := getBoardString(move.ToCol, move.ToRow - 1, gameBoard)

		if (potentialKing == BKing && move.WhitePlayer) || (potentialKing == WKing && !move.WhitePlayer) {
			return true
		}
	}

	if move.ToRow < HEIGHT {
		potentialKing := getBoardString(move.ToCol, move.ToRow + 1, gameBoard)

		if (potentialKing == BKing && move.WhitePlayer) || (potentialKing == WKing && !move.WhitePlayer) {
			return true
		}
	}

	return false
}

func KnightInCheck(move *player.Move, gameBoard *Board) bool {

	//up two and left one
	if move.ToRow - 2 >= 0 && move.ToCol - 1 >= 0 {
		potentialKnight := getBoardString(move.ToCol - 1, move.ToRow - 2, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	//up two and right one
	if move.ToRow - 2 >= 0 && move.ToCol + 1 < WIDTH {
		potentialKnight := getBoardString(move.ToCol + 1, move.ToRow - 2, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	//down two and left one
	if move.ToRow + 2 < HEIGHT && move.ToCol - 1 >= 0 {
		potentialKnight := getBoardString(move.ToCol - 1, move.ToRow + 2, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	//down two and right one
	if move.ToRow + 2 < HEIGHT && move.ToCol - 1 >= 0 {
		potentialKnight := getBoardString(move.ToCol + 1, move.ToRow + 2, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	//left 2 and up one
	if move.ToRow + 1 < HEIGHT && move.ToCol - 2 >= 0 {
		potentialKnight := getBoardString(move.ToCol - 2, move.ToRow + 1, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}
	//left 2 and down one
	if move.ToRow - 1 >= 0 && move.ToCol - 2 >= 0 {
		potentialKnight := getBoardString(move.ToCol - 2, move.ToRow - 1, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	//right 2 and up one
	if move.ToRow + 1 < HEIGHT && move.ToCol + 2 < WIDTH {
		potentialKnight := getBoardString(move.ToCol + 2, move.ToRow + 1, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}
	//right 2 and down one
	if move.ToRow - 1 >= 0 && move.ToCol + 2 < WIDTH {
		potentialKnight := getBoardString(move.ToCol + 2, move.ToRow - 1, gameBoard)

		if (potentialKnight == WKnight && !move.WhitePlayer) || (potentialKnight == BKnight && move.WhitePlayer) {
			return true
		}
	}

	return false
}

func QueenInCheck(move *player.Move, gameBoard *Board) bool {
	//this is just a copy paste of bishop and rook it could later be refactored - I just want it to work for now
	//bishop + rook
	col, row := move.ToCol + 1, move.ToRow - 1

	//right down - zero based
	for col < WIDTH && row >= 0 {
		var potentialQueen = getBoardString(col, row, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}

		col++
		row--
	}

	col, row = move.ToCol - 1, move.ToRow - 1
	//left down --zero based
	for col < WIDTH && row >= 0 {
		var potentialQueen = getBoardString(col, row, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}

		col--
		row--
	}



	//diagonal right left up - zero based
	col, row = move.ToCol + 1, move.ToRow + 1
	//right up
	for col < WIDTH && row >= 0 {
		var potentialQueen = getBoardString(col, row, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}

		col++
		row++
	}

	col, row = move.ToCol - 1, move.ToRow + 1
	//left up zero based
	for col < WIDTH && row >= 0 {
		var potentialQueen = getBoardString(col, row, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}

		col--
		row++
	}

	var up, down, left, right = move.ToRow - 1, move.ToRow + 1, move.ToCol - 1, move.ToCol + 1

	for up - move.ToRow >= 0 {
		var potentialQueen = getBoardString(move.ToCol, up - move.ToRow, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}
	}

	for down + move.ToRow < HEIGHT {
		var potentialQueen = getBoardString(move.ToCol, down + move.ToRow, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}
	}

	for left - move.ToCol >= 0 {
		var potentialQueen = getBoardString(left - move.ToCol, move.ToRow, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}
	}

	for right + move.ToCol < WIDTH {
		var potentialQueen = getBoardString(right + move.ToCol, move.ToRow, gameBoard)

		if (potentialQueen == BQueen && move.WhitePlayer) || (potentialQueen == WQueen && !move.WhitePlayer) {
			return true
		} else {
			if potentialQueen != Empty {
				break
			}
		}
	}

	return false
}