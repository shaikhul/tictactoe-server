package api

import (
	"encoding/json"
	"github.com/chrisfregly/tictactoe"
	"net/http"
)

// TicTacToeHandler defines structure of custom http handler
type TicTacToeHandler struct {
	instance tictactoe.TicTacToe
}

// NewTicTacToeHandler creates a new TicTacToeHandler
func NewTicTacToeHandler(i tictactoe.TicTacToe) *TicTacToeHandler {
	return &TicTacToeHandler{
		instance: i,
	}
}

// ServeHTTP decides which handler to call for given http method and path
func (s *TicTacToeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/game":
		s.getGame(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/game/move":
		s.postGameMove(w, r)
		return
	case r.Method == http.MethodDelete && r.URL.Path == "/game":
		s.deleteGame(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid endpoint"))
	}
}

func (s *TicTacToeHandler) getGame(w http.ResponseWriter, r *http.Request) {
	i := s.instance
	g := SerializeTicTacToe(i)
	j, err := json.Marshal(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (s *TicTacToeHandler) postGameMove(w http.ResponseWriter, r *http.Request) {
	var p PostGameMoveRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.instance.Move(tictactoe.Player(p.Player), p.Row, p.Col)
	if err != nil {
		errRes := PostGameMoveErrorResponse{
			Error: err.Error(),
		}
		b, err := json.Marshal(errRes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	s.getGame(w, r)
}

func (s *TicTacToeHandler) deleteGame(w http.ResponseWriter, r *http.Request) {
	s.instance = tictactoe.NewTicTacToe()
	w.WriteHeader(http.StatusNoContent)
}
