package api

import (
	"github.com/chrisfregly/tictactoe"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSerializeGame(t *testing.T) {
	tt := tictactoe.NewTicTacToe()
	// move X to 0,0
	err := tt.Move(tt.GetTurn(), 0, 0)
	require.NoError(t, err)
	// move O to 2,2
	err = tt.Move(tt.GetTurn(), 2, 2)
	require.NoError(t, err)

	game := SerializeTicTacToe(tt)
	require.Equal(t, string(tt.GetTurn()), game.Turn)
	for i := range game.Board {
		for j := range game.Board[i] {
			if i == 0 && j == 0 {
				require.Equal(t, "X", game.Board[0][0])
			} else if i == 2 && j == 2 {
				require.Equal(t, "O", game.Board[2][2])
			} else {
				require.Empty(t, game.Board[i][j])
			}
		}
	}
	require.Empty(t, game.Winner)
}
