package main

import "os"

func main() {
	args := os.Args

	if len(args) < 2 {
		create_hotseat_game()
		return
	}

	switch args[1] {
	case "--hotseat":
		create_hotseat_game()
		break
	case "--host":
		host_ssh_lobby()
		break
	case "--join":
		join_ssh_lobby()
		break
	case "--help":
	default:
		help()
		break
	}
}
