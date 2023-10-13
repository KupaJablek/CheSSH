# CheSSH
CheSSH is a cli based chess application focused on:
- Hotseat gameplay
- Online gameplay secured with ssh

## Installation
- [Go](https://go.dev/) version 1.18+ required
- [Open SSH](https://www.openssh.com/) required
- [Nerdfont](https://www.nerdfonts.com/) *optional*

Install by cloning this repo:
```
git clone https://github.com/KupaJablek/CheSSH.git
```

## Playing The Game
cd into the directory where you cloned CheSSH and run:
```go
go run ./cmd/CheSSH
```
OR  
run the binary in **/CheSSH/bin/** with:
```
./CheSSH
```

### Commands

Runnable binaries can be found in **/CheSSH/bin/**

To start a **hotseat** game of chess:
```
./CheSSH --hotseat
```

To **host** game of chess over ssh:
```
./CheSSH --host
```

To **join** a game of chess over ssh:
```
./CheSSH --join
```

### Making a Move

During a game moves are made using a modified [algebraic chess notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess)):
```
a1-b3 or C3-C1 or d4-g7
```

## Customization

Customization will be coming in a future update.

## Binaries

### Build

Build the binary with the following make command:
```
make build
```
The resulting binary can be found in **/CheSSH/bin/**

### Clean

Clean **/CheSSH/bin** of **all** binaries with:
```
make clean
```

### Compile

Make binaries for all supported OS:
```
make compile
```
