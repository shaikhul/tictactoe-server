package main

import (
	"github.com/chrisfregly/tictactoe"
	"github.com/shaikhul/tictactoe-server/api"
	"log"
	"net/http"
)

func main() {
	game := tictactoe.NewTicTacToe()
	gameHandler := api.NewTicTacToeHandler(game)

	mux := http.NewServeMux()
	mux.Handle("/game", gameHandler)
	mux.Handle("/game/move", gameHandler)

	log.Panic(http.ListenAndServe(":8080", mux))
}
