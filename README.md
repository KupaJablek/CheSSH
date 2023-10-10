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
go run .
```
OR  
run the binary with:
```
./chessh
```

### Commands

To start a **hotseat** game of chess:
```
./chessh
    OR
./chessh --hotseat
```

To **host** game of chess over ssh:
```
./chessh --host
```

To **join** a game of chess over ssh:
```
./chessh --join
```

### Making a Move

During a game moves are made using a modified [algebraic chess notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess)):
```
a1-b3 or C3-C1 or d4-g7
```

## Customization

Customization will be coming in a future update.