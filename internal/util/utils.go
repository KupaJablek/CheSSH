package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func AbsoluteVal(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Help() {
	fmt.Println("CheSSH is a tool for playing local or online chess")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\tchessh <command> [arguments]")
	fmt.Println("")
	fmt.Println("The commands are:")
	fmt.Println("\t--hotseat")
	fmt.Println("\t--host")
	fmt.Println("\t--join")
	fmt.Println("")
	fmt.Println("To print this menu:")
	fmt.Println("\t--help")
}

func ClearTerminal() {
	platform := runtime.GOOS
	if platform == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
