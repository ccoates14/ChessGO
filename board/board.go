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
	var rowNumber = HEIGHT
	for row := 0; row < HEIGHT; row++ {
		fmt.Print(rowNumber)
		rowNumber--

		for col := 0; col < WIDTH; col++ {
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
	kingX, kingY := currentTeamKingPos(move.WhitePlayer, gameBoard)

	if inCheck(kingX, kingY, move.WhitePlayer, gameBoard) {
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
		return pieceBelongsToWhiteTeam(fromPiece)
	} else {
		return !pieceBelongsToWhiteTeam(fromPiece)
	}
}

func pieceBelongsToWhiteTeam(piece string) bool {
	for _, v := range WhitePieces {
		if piece == v {
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

	return fromPiece != Empty && toPiece != Empty && pieceBelongsToWhiteTeam(fromPiece) == pieceBelongsToWhiteTeam(toPiece)
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

func inCheck(kingX int, kingY int, whitePlayer bool, gameBoard *Board) bool {
	check := false
	if PawnInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("Pawn Check")
		check = true
	} else if KingInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("King Check")
		check = true
	} else if RookInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("Rook Check")
		check = true
	} else if KnightInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("Knight Check")
		check = true
	} else if BishopInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("Bishop Check")
		check = true
	} else if QueenInCheck(kingX, kingY, whitePlayer, gameBoard) {
		fmt.Println("Queen Check")
		check = true
	}

	return check
}

func getBoardString(col int, row int, gameBoard *Board) string {
	return gameBoard.pieces[row][col]
}

func setBoardString(col int, row int, gameBoard *Board, piece string) string {
	currentPiece := getBoardString(col, row, gameBoard)

	gameBoard.pieces[row][col] = piece

	return currentPiece
}

func currentTeamKingPos(isWhiteTeam bool, gameBoard *Board) (int, int) {
	kingX := 0
	kingY := 0

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			currentPiece := getBoardString(i, j, gameBoard)
			if currentPiece == BKing && !isWhiteTeam || currentPiece == WKing && isWhiteTeam {
				kingX = i
				kingY = j
				break
			}
		}
	}

	return kingX, kingY
}