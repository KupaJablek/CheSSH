BINARY_NAME=CheSSH

build:
	go build -o bin/${BINARY_NAME} ./cmd/CheSSH/

hotseat:
	go run ./cmd/CheSSH/ --hotseat
