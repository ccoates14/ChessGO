package board

import (
	"ChessGo/player"
	"fmt"
	"math"
)

func ValidQueenMove(move *player.Move, gameBoard *Board) bool {
	bishopMovement := move.FromCol != move.ToCol && move.FromRow != move.ToRow 

	if bishopMovement {
		validBishop := ValidBishopMove(move, gameBoard)
		fmt.Println("Valid bishop movement for queen: ", validBishop)
		return validBishop
	} else {
		validRook := ValidRookMove(move, gameBoard) 
		fmt.Println("Valid rook move for queen: ", validRook)
		return validRook
	}
}

func ValidPawnMove(move *player.Move, gameBoard *Board) bool {
	//is moving forward
	if move.FromCol == move.ToCol {
		toPiece := getBoardString(move.ToCol, move.ToRow, gameBoard)

		if toPiece != Empty {
			return false
		}

		var difference = math.Abs(float64(move.ToRow - move.FromRow))

		if difference <= 2 {
			if move.WhitePlayer {
				if move.FromRow <= move.ToRow {
					return false
				}
			} else {
				if move.FromRow >= move.ToRow {
					return false
				}
			}
		}

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
		} else {
			if move.WhitePlayer {
				if move.FromRow <= move.ToRow {
					return false
				}
			} else {
				if move.FromRow >= move.ToRow {
					return false
				}
			}
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
	
		var distance = math.Abs(float64(move.FromRow - move.ToRow)) - 1
	
		for i := 1; i < int(distance); i++ {
			//check if current position contains a piece
			if getBoardString(move.FromCol, move.FromRow + (direction * i), gameBoard) != Empty {
				return false
			}
		}

		return true
	} else if move.FromRow == move.ToRow { //left right
		var direction = 1 //right

		if move.FromCol > move.ToCol { // left
			direction = -1
		}
	
		var distance = math.Abs(float64(move.FromCol - move.ToCol)) - 1
	
		for i := 1; i < int(distance); i++ {
			//check if current position contains a piece
			if getBoardString(move.FromCol + (direction * i), move.FromRow, gameBoard) != Empty {
				return false
			}
		}

		return true
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
	if math.Abs(float64(move.FromCol - move.ToCol)) != math.Abs(float64(move.FromRow - move.ToRow)) || move.FromCol == move.ToCol || move.FromRow == move.ToRow {
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

	for i := 1; i < int(distance); i++ {
		if getBoardString(move.FromCol + (h_direction * i), move.FromRow + (v_direction * i), gameBoard) != Empty {
			return false
		}
	}

	return true
}

func PawnInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	//at this point with all these potential in check methods we assume
	//the move has already been performed
	//use the to xy

	verticalPawnDirection := -1

	if !whitePlayer {
		verticalPawnDirection = 1
	}

	if kingY + verticalPawnDirection >= 0 && kingY + verticalPawnDirection < HEIGHT {
		//check left
		if kingX - 1 >= 0 {
			var potentialPawn = getBoardString(kingX - 1, kingY + verticalPawnDirection, gameBoard)
			if (whitePlayer && potentialPawn == BPawn) || (!whitePlayer && potentialPawn == WPawn) {
				return true
			}
		}

		//check right
		if kingX + 1 < WIDTH {
			var potentialPawn = getBoardString(kingX + 1, kingY + verticalPawnDirection, gameBoard)
			if (whitePlayer && potentialPawn == BPawn) || (!whitePlayer && potentialPawn == WPawn) {
				return true
			}
		}
	}

	return false
}

func RookInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	//up down left right

	//starting from king position
	//move out in a cross shape and see if you encounter a rook of the opposing team
	var up, down, left, right = kingY - 1, kingY + 1, kingX - 1, kingX + 1

	for up - kingY >= 0 {
		var potentialRook = getBoardString(kingX, up - kingY, gameBoard)

		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		up--
	}

	for down + kingY < HEIGHT {
		var potentialRook = getBoardString(kingX, down + kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		down++
	}

	for left - kingX >= 0 {
		var potentialRook = getBoardString(left - kingX, kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		left--
	}

	for right + kingX < WIDTH {
		var potentialRook = getBoardString(right + kingX, kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		right++
	}

	return false
}

func BishopInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	col, row := kingX + 1, kingY + 1

	//right down - zero based
	for col < WIDTH && row < HEIGHT {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row++
	}

	col, row = kingX - 1, kingY + 1
	//left down --zero based
	for col >= 0 && row < WIDTH {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row++
	}



	//diagonal right up - zero based
	col, row = kingX + 1, kingY - 1
	//right up
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row--
	}

	col, row = kingX - 1, kingY - 1
	//left up zero based
	for col >= 0 && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row--
	}



	return false
}

func KingInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	if kingX - 1 >= 0 {
		potentialKing := getBoardString(kingX - 1, kingY, gameBoard)

		if (potentialKing == BKing && whitePlayer) || (potentialKing == WKing && !whitePlayer) {
			return true
		}
	}

	if kingX + 1 < WIDTH {
		potentialKing := getBoardString(kingX + 1, kingY, gameBoard)

		if (potentialKing == BKing && whitePlayer) || (potentialKing == WKing && !whitePlayer) {
			return true
		}
	}

	if kingY - 1 >= 0 {
		potentialKing := getBoardString(kingX, kingY - 1, gameBoard)

		if (potentialKing == BKing && whitePlayer) || (potentialKing == WKing && !whitePlayer) {
			return true
		}
	}

	if kingY + 1 < HEIGHT {
		potentialKing := getBoardString(kingX, kingY + 1, gameBoard)

		if (potentialKing == BKing && whitePlayer) || (potentialKing == WKing && !whitePlayer) {
			return true
		}
	}

	return false
}

func KnightInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	//up two and left one
	if kingY - 2 >= 0 && kingX - 1 >= 0 {
		potentialKnight := getBoardString(kingX - 1, kingY - 2, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	//up two and right one
	if kingY - 2 >= 0 && kingX + 1 < WIDTH {
		potentialKnight := getBoardString(kingX + 1, kingY - 2, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	//down two and left one
	if kingY + 2 < HEIGHT && kingX - 1 >= 0 {
		potentialKnight := getBoardString(kingX - 1, kingY + 2, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	//down two and right one
	if kingY + 2 < HEIGHT && kingX - 1 >= 0 {
		potentialKnight := getBoardString(kingX + 1, kingY + 2, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	//left 2 and up one
	if kingY + 1 < HEIGHT && kingX - 2 >= 0 {
		potentialKnight := getBoardString(kingX - 2, kingY + 1, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}
	//left 2 and down one
	if kingY - 1 >= 0 && kingX - 2 >= 0 {
		potentialKnight := getBoardString(kingX - 2, kingY - 1, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	//right 2 and up one
	if kingY + 1 < HEIGHT && kingX + 2 < WIDTH {
		potentialKnight := getBoardString(kingX + 2, kingY + 1, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}
	//right 2 and down one
	if kingY - 1 >= 0 && kingX + 2 < WIDTH {
		potentialKnight := getBoardString(kingX + 2, kingY - 1, gameBoard)

		if (potentialKnight == WKnight && !whitePlayer) || (potentialKnight == BKnight && whitePlayer) {
			return true
		}
	}

	return false
}

func QueenInCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	//this is just a copy paste of bishop and rook it could later be refactored - I just want it to work for now
	//bishop + rook
	col, row := kingX + 1, kingY + 1

	//right down - zero based
	for col < WIDTH && row < HEIGHT {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row++
	}

	col, row = kingX - 1, kingY + 1
	//left down --zero based
	for col >= 0 && row < WIDTH {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row++
	}



	//diagonal right up - zero based
	col, row = kingX + 1, kingY - 1
	//right up
	for col < WIDTH && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col++
		row--
	}

	col, row = kingX - 1, kingY - 1
	//left up zero based
	for col >= 0 && row >= 0 {
		var potentialBishop = getBoardString(col, row, gameBoard)

		if (potentialBishop == BBishop && whitePlayer) || (potentialBishop == WBishop && !whitePlayer) {
			return true
		} else {
			if potentialBishop != Empty {
				break
			}
		}

		col--
		row--
	}

	var up, down, left, right = kingY - 1, kingY + 1, kingX - 1, kingX + 1

	for up - kingY >= 0 {
		var potentialRook = getBoardString(kingX, up - kingY, gameBoard)

		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		up--
	}

	for down + kingY < HEIGHT {
		var potentialRook = getBoardString(kingX, down + kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		down++
	}

	for left - kingX >= 0 {
		var potentialRook = getBoardString(left - kingX, kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		left--
	}

	for right + kingX < WIDTH {
		var potentialRook = getBoardString(right + kingX, kingY, gameBoard)
		if (potentialRook == BRook && whitePlayer) || (potentialRook == WRook && !whitePlayer) {
			return true
		} else {
			if potentialRook != Empty {
				break
			}
		}

		right++
	}

	return false
}