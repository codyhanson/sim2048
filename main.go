package main

import (
	"fmt"
	"github.com/codyhanson/sim2048/game"
	"os"
	"os/exec"
)
import "github.com/pkg/term"

func main() {

	cmd := exec.Command("clear") // sorry windows.
	cmd.Stdout = os.Stdout
	cmd.Run()

	board := game.New()

	for {
		board.Print()
		command := nextCommand()

		if command == game.QUIT {
			break
		}

		board.Shift(command)
		if board.BlockSameMove != command {
			//only gen squares if we aren't on a blocked move
			board.GenNewSquares()
		}
		if board.Won {
			fmt.Println("You won!")
			break
		}

		if board.Lost {
			fmt.Println("You lost!")
			break
		}

	}
}

func nextCommand() game.Command {
	ascii, keycode, err := getChar()
	if err != nil {

	}
	if keycode == 38 {
		return game.UP
	}
	if keycode == 40 {
		return game.DOWN
	}
	if keycode == 37 {
		return game.LEFT
	}
	if keycode == 39 {
		return game.RIGHT
	}
	if ascii == 99 {
		return game.QUIT
	}
	return game.QUIT
}

// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
