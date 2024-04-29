BINARY_NAME=CheSSH

build:
	go build -o bin/${BINARY_NAME} ./cmd/CheSSH/

hotseat:
	go run ./cmd/CheSSH/ --hotseat

clean:
	go clean
	rm ./bin/*

tests:
	go test -v ./internal/game/util.go ./internal/game/game.go ./internal/game/game_test.go

help:
	@printf "usage: make <command>\n\n"
	@printf "commands are:\n\n"
	
	@printf "	build: build a binary for your system\n"
	@printf "	clean: remove all binaries from ./bin/\n"
	@printf "	tests: run all tests\n\n"
	@printf "	help: show this menu\n"
