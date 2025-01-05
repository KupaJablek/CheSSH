# CheSSH

CheSSH is a cli chess application focused on hotseat chess gameplay

## Installation
Make sure to have the following dependencies installed:

1. [Go](https://go.dev/) version 1.18+ required
2. [Open SSH](https://www.openssh.com/) required

Clone this repo:
```
git clone https://github.com/KupaJablek/CheSSH.git
```

From your install dir, run:
```
make build
```

To run from anywhere add this command to your .bashrc:
```
alias CheSSH='~/*your install path*/CheSSH/bin/CheSSH'
```

## Playing the game

To run the application use `CheSSH` + one of the params below:

### Hotseat

Use `--hotseat` to run a local hotseat session for you and one other player

### Online via SSH

Use `--Host` to host an online lobby for another player to join

OR

Use `--Join` to join another player's lobby

Joining and hosting can be configured with the `--ip`, `--roomname`, or `--port` flags. 

### Making a Move

During a game moves are made using a modified [algebraic chess notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess)):
```
a1-b3 or c3-c1 or d4-g7
```

## Makefile

For a full list of commands run `make help`

- You can build your installation with `make build`

- Delete binaries with `make clean`

## License
[mit](LICENSE)