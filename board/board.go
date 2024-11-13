package board

import (
	"ChessGo/player"
	"fmt"
	"errors"
)

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

var WhitePieces = [6]string{WPawn, WRook, WKing, WKnight, WQueen, WBishop}

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
			fmt.Print(" " + getBoardString(col, row, gameBoard) + " ")
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

//make this return the piece that was killed or empty if empty
func AttemptMove(move *player.Move, gameBoard *Board) error {

	if !moveWithinBoard(move) {
		return errors.New("Move not within Board")
	} 

	if !pieceBelongsToPlayer(move, gameBoard) {
		return errors.New("Piece Belongs to Other Player or move from position is empty")
	} 

	if moveAttackingSameTeam(move, gameBoard) {
		return errors.New("Move attacking same Team")
	}

	if !validPieceMove(move, gameBoard) {
		return errors.New("Not a valid move for the Piece")
	} 

	//if it reaches this point we will perform the move
	toPiece := setBoardString(move.ToCol, move.ToRow, gameBoard, getBoardString(move.FromCol, move.FromRow, gameBoard))
	
	//empty from Pos
	currentPiece := setBoardString(move.FromCol, move.FromRow, gameBoard, Empty)

	if inCheck(move, gameBoard) {
		//undo the move
		setBoardString(move.FromCol, move.FromRow, gameBoard, currentPiece)
		setBoardString(move.ToCol, move.ToRow, gameBoard, toPiece)
		
		return errors.New("move puts in check")
	} 
	

	return nil
}

//also will check that it is not an empty piece being grabbed
func pieceBelongsToPlayer(move *player.Move, gameBoard *Board) bool {
	var fromPiece = getBoardString(move.FromCol, move.FromRow, gameBoard)

	if fromPiece == Empty {
		return false
	}

	if move.WhitePlayer {
		return pieceBelongsToWhiteTeam(&fromPiece)
	} else {
		return !pieceBelongsToWhiteTeam(&fromPiece)
	}
}

func pieceBelongsToWhiteTeam(piece *string) bool {
	for _, v := range WhitePieces {
		if *piece == v {
			return true
		}
	}

	return false
}

func moveWithinBoard(move *player.Move) bool {
	return move.FromCol >= 0 && move.FromCol < HEIGHT &&
		move.FromRow >= 0 && move.FromRow < WIDTH &&
		move.ToCol >= 0 && move.ToCol < HEIGHT &&
		move.ToRow >= 0 && move.ToRow < WIDTH
}

func moveAttackingSameTeam(move *player.Move, gameBoard *Board) bool {
	var fromPiece = getBoardString(move.FromCol, move.FromRow, gameBoard)
	var toPiece = getBoardString(move.ToCol, move.ToRow, gameBoard)

	return pieceBelongsToWhiteTeam(&fromPiece) == pieceBelongsToWhiteTeam(&toPiece)
}

func validPieceMove(move *player.Move, gameBoard *Board) bool {
	var fromPiece = getBoardString(move.FromCol, move.FromRow, gameBoard)

	switch fromPiece {
		case WBishop, BBishop:
			return ValidBishopMove(move, gameBoard)
		case WKing, BKing:
			return ValidKingMove(move, gameBoard)
		case WPawn, BPawn:
			return ValidPawnMove(move, gameBoard)
		case WRook, BRook:
			return ValidRookMove(move, gameBoard)
		case WQueen, BQueen:
			return ValidQueenMove(move, gameBoard)
		case WKnight, BKnight:
			return ValidKnightMove(move, gameBoard)
		default:
			fmt.Println("Error: Unknown game piece for valid move check")
			return false
	}
}

func inCheck(move *player.Move, gameBoard *Board) bool {
	return PawnInCheck(move, gameBoard) ||
		QueenInCheck(move, gameBoard) ||
		KingInCheck(move, gameBoard) || 
		RookInCheck(move, gameBoard) || 
		KnightInCheck(move, gameBoard) ||
		BishopInCheck(move, gameBoard)
}

func getBoardString(col int, row int, gameBoard *Board) string {
	return gameBoard.pieces[row][col]
}

func setBoardString(col int, row int, gameBoard *Board, piece string) string {
	currentPiece := getBoardString(col, row, gameBoard)

	gameBoard.pieces[row][col] = piece

	return currentPiece
}