package main

import (
	"bytes"
	"encoding/json"
	"github.com/chrisfregly/tictactoe"
	"github.com/shaikhul/tictactoe-server/api"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGameHandler_GetGameEndpoint(t *testing.T) {
	i := tictactoe.NewTicTacToe()
	s := api.NewTicTacToeHandler(i)

	// GET - get current game state
	r := httptest.NewRequest(http.MethodGet, "/game", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, 200, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	var gameRes api.GameResponse
	err = json.Unmarshal(data, &gameRes)
	require.NoError(t, err)
	require.Equal(t, "X", gameRes.Turn)
	require.Empty(t, gameRes.Winner)
}

func TestGameHandler_PostGameMoveEndpoint(t *testing.T) {
	i := tictactoe.NewTicTacToe()
	s := api.NewTicTacToeHandler(i)

	// POST - make a move for X
	rawJson := `{"player": "X", "row": 0, "col": 0}`
	jsonBytes := []byte(rawJson)

	r := httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, 200, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	var gameRes api.GameResponse
	err = json.Unmarshal(data, &gameRes)
	require.NoError(t, err)
	require.Equal(t, "O", gameRes.Turn)
	require.Empty(t, gameRes.Winner)
	require.Equal(t, "X", gameRes.Board[0][0])

	// POST - make a move for O
	rawJson = `{"player": "O", "row": 0, "col": 1}`
	jsonBytes = []byte(rawJson)

	r = httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res = w.Result()
	defer res.Body.Close()
	require.Equal(t, 200, res.StatusCode)

	data, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	err = json.Unmarshal(data, &gameRes)
	require.NoError(t, err)
	require.Equal(t, "X", gameRes.Turn)
	require.Empty(t, gameRes.Winner)
	require.Equal(t, "X", gameRes.Board[0][0])
	require.Equal(t, "O", gameRes.Board[0][1])

	// POST - make an invalid move, should get "tictactoe: not Z's turn"
	rawJson = `{"player": "Z", "row": 0, "col": 0}`
	jsonBytes = []byte(rawJson)

	r = httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res = w.Result()
	defer res.Body.Close()
	require.Equal(t, 200, res.StatusCode)

	data, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	var errRes api.PostGameMoveErrorResponse
	err = json.Unmarshal(data, &errRes)
	require.NoError(t, err)
	require.Equal(t, "tictactoe: not Z's turn", errRes.Error)

	// POST - get a winner
	rawJson = `{"player": "X", "row": 1, "col": 1}`
	jsonBytes = []byte(rawJson)

	r = httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)

	rawJson = `{"player": "O", "row": 1, "col": 0}`
	jsonBytes = []byte(rawJson)

	r = httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)

	rawJson = `{"player": "X", "row": 2, "col": 2}`
	jsonBytes = []byte(rawJson)

	r = httptest.NewRequest(http.MethodPost, "/game/move", bytes.NewReader(jsonBytes))
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res = w.Result()
	defer res.Body.Close()
	require.Equal(t, 200, res.StatusCode)

	data, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	err = json.Unmarshal(data, &gameRes)
	require.NoError(t, err)
	require.Equal(t, "X", gameRes.Turn)
	require.Equal(t, "X", gameRes.Winner)
	require.Equal(t, "X", gameRes.Board[0][0])
	require.Equal(t, "X", gameRes.Board[1][1])
	require.Equal(t, "X", gameRes.Board[2][2])
}

func TestGameHandler_DeleteGameEndpoint(t *testing.T) {
	i := tictactoe.NewTicTacToe()
	s := api.NewTicTacToeHandler(i)

	// DELETE - reset current game
	r := httptest.NewRequest(http.MethodDelete, "/game", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, 204, res.StatusCode)
}
