package board

import "fmt"
import "ChessGo/player"

//board will be 2d array of characters representing Chess 
const (
    Empty    string = "·"
    WPawn    string = "♙"
    WRook    string = "♖"
    WKnight  string = "♘"
    WBishop  string = "♗"
    WQueen   string = "♕"
    WKing    string = "♔"
    BPawn    string = "♟"
    BRook    string = "♜"
    BKnight  string = "♞"
    BBishop  string = "♝"
    BQueen   string = "♛"
    BKing    string = "♚"
	WIDTH	 int   =  8
	HEIGHT   int   =  8
)

type Board struct {
	pieces [HEIGHT][WIDTH]string
}

func InitializeBoard() *Board {
    return &Board {
		[HEIGHT][WIDTH]string{
			{BRook, BKnight, BBishop, BQueen, BKing, BBishop, BKnight, BRook},
			{BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn},
			{WRook, WKnight, WBishop, WQueen, WKing, WBishop, WKnight, WRook},
    	},
	}
}

func RenderBoard(gameBoard *Board) {
	var colNumber = HEIGHT
	for col := 0; col < HEIGHT; col++ {
		fmt.Print(colNumber)
		colNumber--

		for row := 0; row < WIDTH; row++ {
			fmt.Print(" " + getBoardstring(col, row, gameBoard) + " ")
		}
		fmt.Print("\n")
	}

	fmt.Print("  ")

	for rowNumber := 65; rowNumber < 65 + WIDTH; rowNumber++ {
		fmt.Printf("%c ", rowNumber)
		fmt.Print(" ")
	}
	fmt.Print("\n")

}

// func GetWinner() Player or nil
	//if there is a checkmate that team loses
	//if there is a certain number of consecutive checks then there is a tie 

// func AttemptMove() returns true or false if move was executed - false meaning it was an invalid move
	//is move within board?
		//is move not attacking same team?
			//can the current piece move to the position being requested?
				//would this move put the current team into check?
func AttemptMove(move *player.Move, gameBoard *Board) *string {
	var errorMessage string = "Error: "

	if moveWithinBoard(move, gameBoard) {
		if pieceBelongsToPlayer(move, gameBoard) {
			if !moveAttackingSameTeam(move, gameBoard) {
				if validPieceMove(move, gameBoard) {
					if !potentialCheck(move, gameBoard) {
						//perform the actual move
					} else {
						errorMessage += "Would put you into check"
					}
				} else {
					errorMessage += "Not a valid move for the Piece"
				}
			} else {
				errorMessage += "Move attacking same Team"
			}
		} else {
			errorMessage += "Piece Belongs to Other Player or move from position is empty" 
		}
	} else {
		errorMessage += "Move not within Board"
	}

	return &errorMessage
}

func pieceBelongsToPlayer(move *player.Move, gameBoard *Board) bool {
	return false
}

func moveWithinBoard(move *player.Move, gameBoard *Board) bool {
	return false
}

func moveAttackingSameTeam(move *player.Move, gameBoard *Board) bool {
	return false
}

func validPieceMove(move *player.Move, gameBoard *Board) bool {
	return false
}

func potentialCheck(move *player.Move, gameBoard *Board) bool {
	return false
}

func getBoardstring(col int, row int, gameBoard *Board) string {
	return gameBoard.pieces[col][row]
}