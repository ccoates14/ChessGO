package main

import (
	"ChessGo/board"
	"ChessGo/player"
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
	 "unicode/utf8"
)

func main() {
    fmt.Println(`          _             _       _    _           _           _                  _          _                    _              _          _     
        /\ \           / /\    / /\ /\ \        / /\        / /\               /\ \       /\ \     _           /\ \           /\ \       /\_\   
       /  \ \         / / /   / / //  \ \      / /  \      / /  \              \ \ \     /  \ \   /\_\        /  \ \         /  \ \     / / /   
      / /\ \ \       / /_/   / / // /\ \ \    / / /\ \__  / / /\ \__           /\ \_\   / /\ \ \_/ / /       / /\ \_\       / /\ \ \   / / /_   
     / / /\ \ \     / /\ \__/ / // / /\ \_\  / / /\ \___\/ / /\ \___\         / /\/_/  / / /\ \___/ /       / / /\/_/      / / /\ \ \ / /___/\  
    / / /  \ \_\   / /\ \___\/ // /_/_ \/_/  \ \ \ \/___/\ \ \ \/___/        / / /    / / /  \/____/       / / / ______   / / /  \ \_\\____ \ \ 
   / / /    \/_/  / / /\/___/ // /____/\      \ \ \       \ \ \             / / /    / / /    / / /       / / / /\_____\ / / /   / / /    / / / 
  / / /          / / /   / / // /\____\/  _    \ \ \  _    \ \ \           / / /    / / /    / / /       / / /  \/____ // / /   / / /    / / /  
 / / /________  / / /   / / // / /______ /_/\__/ / / /_/\__/ / /       ___/ / /__  / / /    / / /       / / /_____/ / // / /___/ / /    _\/_/   
/ / /_________\/ / /   / / // / /_______\\ \/___/ /  \ \/___/ /       /\__\/_/___\/ / /    / / /       / / /______\/ // / /____\/ /    /\_\     
\/____________/\/_/    \/_/ \/__________/ \_____\/    \_____\/        \/_________/\/_/     \/_/        \/___________/ \/_________/     \/_/     
                                                                                                                                                `)

	//
	// This will be the complete game of Chess, including two player more and a really basic AI to play against. 
	// The first version will be purely in cmd and depending on how long this takes me I will come back around later and add an actual UI.
	// The default will be via cmd but a flag will later be added to launch with UI.
	// 

	// TODO:

	// 1. Build basic game loop with IO
	// 2. 

	// structure

	// main
	//	game loop

	// player

	// ai 

	// board
	
	// piece

	fmt.Print("Enter moves fromCol,fromRow, toCol, toRow\n")
	fmt.Print("For example, 1B 3A\n")
	fmt.Print("This would move the white knight to the position 3A\n\n")


	board := board.InitializeBoard()

	gameLoop(board)

}

func gameLoop(gameBoard *board.Board) {
	reader := bufio.NewReader(os.Stdin)
	var move string
	whiteMove := true
	winnerIsWhiteTeam := true
	done := false

	for !done {
		//display board
		board.RenderBoard(gameBoard)

		//ask player for move
		fmt.Println()
		if whiteMove {
			fmt.Print("White Team")
		} else {
			fmt.Print("Black Team")
		}

		fmt.Print(" enter your move: ")
		move, _ = reader.ReadString('\n')

		move = strings.Trim(move, "\n")

		parsedMoved, err := parseMove(move)

		if err != nil {
			fmt.Println(err)
		} else {
			parsedMoved.WhitePlayer = whiteMove

			boardError := board.AttemptMove(&parsedMoved, gameBoard)

			if boardError != nil {
				fmt.Println(boardError)
			} else {
				done, winnerIsWhiteTeam = gameOver()
				whiteMove = !whiteMove
			}
		}
	}

	if winnerIsWhiteTeam {
		fmt.Println("Winner is white team")
	} else {
		fmt.Println("Winner is black team")
	}
}

func parseMove(move string) (player.Move, error) {
	errorMessage := "Invalid move - must be XX XX, such as A8 B6"
	moveStruct := player.Move{}

	if utf8.RuneCountInString(move) != 5 {
		return moveStruct, errors.New(errorMessage + ", string to long: " + strconv.Itoa(utf8.RuneCountInString(move)))
	}

	firstMove := strings.Split(move, " ")[0]
	secondMove := strings.Split(move, " ")[1]

	if !validMovePart(firstMove) || !validMovePart(secondMove) {
		return moveStruct, errors.New(errorMessage)
	}

	//convert the string numbers into actual board numbers
	moveStruct.FromCol = int(math.Abs(float64(int('A') - int(firstMove[0]))))
	moveStruct.FromRow = int(math.Abs(math.Abs(float64(int('1') - int(firstMove[1]))) - float64(board.HEIGHT - 1)))


	moveStruct.ToCol = int(math.Abs(float64(int('A') - int(secondMove[0]))))
	moveStruct.ToRow = int(math.Abs(math.Abs(float64(int('1') - int(secondMove[1]))) - float64(board.HEIGHT - 1)))

	return moveStruct, nil
}

func validMovePart(movePart string) bool {
	return movePart[0] >= 'A' && movePart[0] <= 'H' && movePart[1] >= '1' && movePart[1] <= '8'
}

func gameOver() (bool, bool) {
	return false, false
}