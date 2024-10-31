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