package game

import "fmt"

func printGameStats(p1w, p2w int, p1n, p2n string) {
    fmt.Printf("%s | Wins: %d | Losses: %d\n", p1n, p1w, p2w)
    fmt.Printf("%s | Wins: %d | Losses: %d\n", p2n, p2w, p1w)
}

func HostHotseat() {
    var p1wins, p2wins int
    p1name := "Player 1"
    p2name := "Player 2"
    // load and store config

    for {
        var game Game
        InitializeBoard(&game)
        game.p1Turn = true 
        for !game.gameOver {
            PrintBoard(&game)
            TakeTurn(&game)
        }

        if game.p1Winner {
            p1wins++
        } else {
            p2wins++
        }

        fmt.Println()
        printGameStats(p1wins, p2wins, p1name, p2name)
        var opt string
        fmt.Println("Rematch? y/n:")
        fmt.Scanln(&opt)
        switch opt {
        case "y":
        case "n", "N":
            fmt.Println("Exiting the program")
            return
        default:
            fmt.Println("Exiting the program")
            return
        }
    }
}
