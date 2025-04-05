package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the terminal scren for better view

func ClearScreen() {
	// rutnime.GOOS help to find the target operating system name at compile time
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

	default:
		fmt.Print("\033[2J\033[H")

	}
}
