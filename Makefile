BINARY_NAME=CheSSH

build:
	go build -o bin/${BINARY_NAME} ./cmd/CheSSH/

clean:
	go clean
	rm ./bin/*
run:
	go run ./cmd/CheSSH/

hotseat:
	go run ./cmd/CheSSH/ --hotseat

compile:
	echo "Compiling ${BINARY_NAME} for all supported OS"
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}_linux_x64 ./cmd/CheSSH/
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}_darwin_x64 ./cmd/CheSSH/
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}_windows_x64 ./cmd/CheSSH/
	GOOS=freebsd GOARCH=amd64 go build -o bin/${BINARY_NAME}_freebsd_x64 ./cmd/CheSSH/