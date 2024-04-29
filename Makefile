BINARY_NAME=CheSSH

build:
	go build -o bin/${BINARY_NAME} ./cmd/CheSSH/

hotseat:
	go run ./cmd/CheSSH/ --hotseat

tests:
	go test -v ./internal/game/util.go ./internal/game/game.go ./internal/game/game_test.go
