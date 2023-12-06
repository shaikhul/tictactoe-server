## TicTacToe webserver implementation

### Setup 
#### Docker
* Build docker image: `$ docker build -t tictactoe-server .`
* Run docker image as container exposing server at port `8080`: `$ docker run --publish 8080:8080 tictactoe-server`
* Run tests: `docker run tictactoe-server go test -v ./...`
* 
#### local development
You need golang v1.20.
* Download dependencies: `go mod download`
* Build binary: `go build -o ./bin/tictactoe-server cmd/main.go`
* Run server: `./bin/tictactoe-server`
* Run tests: `$  go test -v ./...`

### API Schema 
* `GET /game` endpoint sample json response 
```
{
  "turn": "X",
  "board": [
    [
      "",
      "",
      ""
    ],
    [
      "",
      "",
      ""
    ],
    [
      "",
      "",
      ""
    ]
  ]
}
```

* `POST /game/move` endpoint requires json data as follows
```
{
"player": "X",
"row": 0,
"col": 0
}
```

* `POST /game/move` endpoint error response
```
{
  "error": "tictactoe: location 0,0 is not empty"
}
```

### Assumptions
* Only two player with name `X` and `O` can play the game
* Board size is 3x3
* The server starts with an empty state, and `X`'s turn
* Once the game is over, server doesn't reset the game unless it gets `DELETE /game` call
* I have used `X`, `O` or `""` to represent a state of each board cell.


### Test Results
```
$ go test -v ./...
=== RUN   TestSerializeGame
--- PASS: TestSerializeGame (0.00s)
PASS
ok      github.com/shaikhul/tictactoe-server/api        (cached)
=== RUN   TestGameHandler_GetGameEndpoint
--- PASS: TestGameHandler_GetGameEndpoint (0.00s)
=== RUN   TestGameHandler_PostGameMoveEndpoint
--- PASS: TestGameHandler_PostGameMoveEndpoint (0.00s)
=== RUN   TestGameHandler_DeleteGameEndpoint
--- PASS: TestGameHandler_DeleteGameEndpoint (0.00s)
PASS
ok      github.com/shaikhul/tictactoe-server/cmd        (cached)
```

### Library used
* Go standard library `net/http`, 
* For tests assertion, I have used third party library `stretchr/testify/require`
