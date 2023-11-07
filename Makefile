BINARY_NAME=CheSSH

build: ##- build a binary for your system
	go build -o bin/${BINARY_NAME} ./cmd/CheSSH/

clean: ##- run go clean and remove all binaries from /bin
	go clean
	rm ./bin/*

test:
	go test -v internal/online/server_utils.go internal/online/online_test.go
	go test -v cmd/CheSSH/main.go cmd/CheSSH/main_test.go

compile: ##- build a binary for all supported x64 OS
	echo "Compiling ${BINARY_NAME} for all supported OS"
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}_linux_x64 ./cmd/CheSSH/
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}_darwin_x64 ./cmd/CheSSH/
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}_windows_x64 ./cmd/CheSSH/
	GOOS=freebsd GOARCH=amd64 go build -o bin/${BINARY_NAME}_freebsd_x64 ./cmd/CheSSH/

help: ##- show this help.
	@printf "\nusage: make <command>\n\n"
	@printf "commands are:\n\n"
	
	@printf "	build: build a binary for your system\n"
	@printf "	clean: remove all binaries from ./bin/\n"
	@printf "	run: run CheSSH with no parameters\n"
	@printf "	hotseat: create a hotseat game with no parameters\n"
	@printf "	compile: create a binary for all supported x64 os\n\n"
	@printf "	help: show this menu\n"
