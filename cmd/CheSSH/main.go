package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args

    if len(args) != 2 {
        fmt.Println("invalid commands")
        return
    }

    switch args[1] {
    case "--hotseat":
        fmt.Println("hotseat")
    case "--join":
        fmt.Println("join")
    case "--host":
        fmt.Println("host")
    }
    return
}
