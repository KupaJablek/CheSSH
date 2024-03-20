package main

import (
	"fmt"
	"os"

	"github.com/KupaJablek/CheSSH/internal/game"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("invalid Args")
        showHelp()
        return
    }

    switch args[1] {
    case "--join":
        game.Join()
    case "--host":
        game.Host()
    case "--hotseat":
        game.HostHotseat()
    default:
        showHelp()
    }
}

func showHelp() {
    fmt.Println("Usage: CheSSH --[hotseat/join/host]")
}

