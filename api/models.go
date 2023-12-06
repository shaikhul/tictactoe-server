package api

// GameResponse defines api response structure for GET and POST method
type GameResponse struct {
	Turn   string     `json:"turn"`
	Winner string     `json:"winner,omitempty"`
	Board  [][]string `json:"board"`
}

// PostGameMoveRequest defines json request structure for POST method
type PostGameMoveRequest struct {
	Player string `json:"player"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
}

// PostGameMoveErrorResponse defines structure for invalid POST response
type PostGameMoveErrorResponse struct {
	Error string `json:"error"`
}

// InitGameResponse initialize GameResponse with empty board
func InitGameResponse(size int) GameResponse {
	r := GameResponse{}
	r.Board = make([][]string, size)
	for i := 0; i < size; i++ {
		r.Board[i] = make([]string, size)
	}
	return r
}
