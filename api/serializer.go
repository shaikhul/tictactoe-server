package api

import (
	"github.com/chrisfregly/tictactoe"
)

// SerializeTicTacToe serializes TicTacToe into a GameResponse
func SerializeTicTacToe(t tictactoe.TicTacToe) GameResponse {
	board := t.GetBoard()
	game := InitGameResponse(len(board))
	game.Turn = string(t.GetTurn())
	if t.GetWinner() != nil {
		game.Winner = string(*t.GetWinner())
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] != nil {
				game.Board[i][j] = string(*board[i][j])
			}
		}
	}
	return game
}
