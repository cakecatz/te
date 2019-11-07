package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unicode"

	"github.com/pkg/term/termios"
)

func refreshScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var origTermios syscall.Termios

func disableRawMode() {
	termios.Tcsetattr(os.Stdin.Fd(), termios.TCSAFLUSH, &origTermios)
}

func enableRawMode() {
	termios.Tcgetattr(os.Stdin.Fd(), &origTermios)

	raw := origTermios
	raw.Lflag &= ^uint32(syscall.ECHO | syscall.ICANON)
	termios.Tcsetattr(os.Stdin.Fd(), termios.TCSAFLUSH, &raw)
}

func main() {
	enableRawMode()
	refreshScreen()

	reader := bufio.NewReader(os.Stdin)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		if r == 'q' {
			break
		}

		if unicode.IsControl(r) {
			fmt.Printf("%d", r)
		} else {
			fmt.Printf("%c", r)
		}

	}

	disableRawMode()
}
