package game

import (
	"fmt"
	"strconv"
)

type Square struct {
	Value int
	AlreadyCombined bool
}

func (square *Square) String() string {
	return fmt.Sprintf("%4s", strconv.Itoa(square.Value))
}