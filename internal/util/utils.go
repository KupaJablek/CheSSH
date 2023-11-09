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
	P1Colour    int
	P2Colour    int
	BoardColour int

	King   string
	Knight string
	Queen  string
	Bishop string
	Rook   string
	Pawn   string

	SshKey     string
	KnownHosts string
}

// initialize default
func InitDefault(c *Config) {
	if c.P1Colour == 0 {
		c.P1Colour = 34
	}

	if c.P2Colour == 0 {
		c.P2Colour = 31
	}

	if c.BoardColour == 0 {
		c.BoardColour = 37
	}

	if c.King == "" {
		c.King = "K"
	}

	if c.Queen == "" {
		c.Queen = "Q"
	}

	if c.Knight == "" {
		c.Knight = "H"
	}

	if c.Bishop == "" {
		c.Bishop = "B"
	}

	if c.Rook == "" {
		c.Rook = "R"
	}

	if c.Pawn == "" {
		c.Pawn = "p"
	}
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
