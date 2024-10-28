package main

import (
	"bufio"
    "fmt"
    "os"
)

import "ChessGo/board"

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

	fmt.Print("Enter moves fromCol,fromRow,toCol,toRow\n")
	fmt.Print("For example, 1,B,3,A\n")
	fmt.Print("This would move the white knight to the position 3A\n\n")


	board := board.InitializeBoard()

	gameLoop(board)

}

func gameLoop(gameBoard *board.Board) {
	reader := bufio.NewReader(os.Stdin)
	var move string

	for !gameOver() {
		//display board
		board.RenderBoard(gameBoard)

		//ask player for move
		fmt.Print("\nEnter your move: ")
		move, _ = reader.ReadString('\n')
		fmt.Print(move)

		//check legal move

		//perform move if legal

		//else tell player illegal and ask again
	}
}

func gameOver() bool {
	return false
}