package board

import "fmt"

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
	strings [8][8]string
}

func InitializeBoard() *Board {
    return &Board {
		[8][8]string{
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
	for col := 0; col < HEIGHT; col++ {
		for row := 0; row < WIDTH; row++ {
			fmt.Print(getBoardstring(col, row, gameBoard) + " ")
		}
		fmt.Print("\n")
	}

}

// func GetWinner() Player or nil
	//if there is a checkmate that team loses
	//if there is a certain number of consecutive checks then there is a tie 

// func AttemptMove() returns true or false if move was executed - false meaning it was an invalid move

func getBoardstring(col int, row int, gameBoard *Board) string {
	return gameBoard.strings[col][row]
}