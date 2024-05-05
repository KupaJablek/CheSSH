package game

import (
    "runtime"
    "os/user"
)

func Abs(i int) int {
    if i < 0 {
        return i * -1
    }
    return i
}

func currentUserKeyPath() string {
	platform := runtime.GOOS
	if platform == "Windows" {
		// not yet implemented
		return ""
	} else {
		user, _ := user.Current()
		return "/home/" + user.Username + "/.ssh/id_rsa"
	}
}

func knownHostsPath() string {
    platform := runtime.GOOS
    if platform == "Windows" {
        // not yet implemented
        return ""
    } else {
        user, _ := user.Current()
        return "/home/" + user.Username + "/.ssh/known_hosts"
    }
}
