package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/BurntSushi/toml"
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

type Config struct {
	P1Colour    string
	P2Colour    string
	BoardColour string

	Knight string
	Queen  string
	Bishop string
	Rook   string
	Pawn   string

	SshKey     string
	KnownHosts string
}

func LoadConfig() (Config, error) {
	var conf Config
	var configPath string
	user, _ := user.Current()
	platform := runtime.GOOS

	if platform == "windows" {
		configPath = ""
		return conf, errors.New("WINDOWS NOT SUPPORTED... yet")
	} else {
		configPath = "/home/" + user.Username + "/.config/CheSSH/config.toml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return conf, errors.New("Cannot read config @: " + configPath + " ERROR: " + err.Error())
	}

	_, err = toml.Decode(string(data), &conf)
	if err != nil {
		return conf, errors.New("Error decoding config @: " + configPath + " ERROR: " + err.Error())
	}

	return conf, nil
}
