package game

type Command int

const (
	NOOP Command = iota
	UP
	DOWN
	LEFT
	RIGHT
	QUIT
)
