package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unicode"

	"github.com/pkg/term/termios"
)

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

		} else {
			fmt.Printf("%c", r)
		}

	}

	disableRawMode()
}
