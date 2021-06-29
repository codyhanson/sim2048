package game

import (
	"fmt"
	"math/rand"
	"time"
)

//0  1  2  3
//4  5  6  7
//8  9  10 11
//12 13 14 15

type Board struct {
	Grid          [16]Square
	BlockSameMove Command
	Won           bool
	Lost          bool
	MoveCount     int
	seed          int64
}

func New() *Board {

	seed := time.Now().UTC().UnixNano()
	//fmt.Printf("Game seed:%d\n", seed)

	board := &Board{
		Grid: [16]Square{},
		seed: seed,
	}

	rand.Seed(seed)

	//start with 2 new squares
	board.GenNewSquares()
	board.GenNewSquares()
	return board
}

func (board *Board) GenNewSquares() {
	// randomly place a new 2
	indexesWithZeros := make([]int, 0, 0)
	for i, v := range board.Grid {
		if v.Value == 0 {
			indexesWithZeros = append(indexesWithZeros, i)
		}
	}

	if len(indexesWithZeros) == 0 {
		board.Lost = true
		return
	}

	i := rand.Intn(len(indexesWithZeros))
	board.setSquareValue(indexesWithZeros[i], 2)

}

func (board *Board) Shift(direction Command) {
	board.MoveCount++
	//TODO if no movement, don't allow same move again.
	if direction == UP && board.BlockSameMove != UP {
		board.iterateForUp()
	} else if direction == DOWN && board.BlockSameMove != DOWN {
		board.iterateForDown()
	} else if direction == LEFT && board.BlockSameMove != LEFT {
		board.iterateForLeft()
	} else if direction == RIGHT && board.BlockSameMove != RIGHT {
		board.iterateForRight()
	}
	board.AdvanceGrid()
}

func (board *Board) AdvanceGrid() {
	for i, _ := range board.Grid {
		board.Grid[i].AlreadyCombined = false
	}
}

func (board *Board) iterateForDown() {
	for i := 15; i >= 12; i-- {
		initial := i
		for j := i; j >= initial-12; j -= 4 {
			board.shiftDown(j)
		}
	}

}

func (board *Board) iterateForLeft() {
	for i := 0; i < 16; i++ {
		board.shiftLeft(i)
	}

}

func (board *Board) iterateForRight() {
	for i := 0; i <= 12; i += 4 {
		initial := i
		for j := initial + 3; j >= initial; j-- {
			board.shiftRight(j)
		}
	}

}

func (board *Board) iterateForUp() {
	for i := 15; i >= 12; i-- {
		initial := i
		for j := initial - 12; j <= i; j += 4 {
			board.shiftUp(j)
		}
	}
}

func (board *Board) shiftUp(nextIndex int) {
	tilesMoved := false
	for {
		prev := nextIndex
		thisCell := board.getSquare(prev)
		if thisCell.Value == 0 {
			return
		}
		nextIndex = nextIndex - 4
		if nextIndex < 0 {
			//already at the edge
			break
		}
		next := board.getSquare(nextIndex)
		if next.Value == 0 {
			//sliiiide it.
			board.setSquareValue(nextIndex, thisCell.Value)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
		} else if next.Value == thisCell.Value && !next.AlreadyCombined && !thisCell.AlreadyCombined {
			//combine
			board.combineSquare(nextIndex)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
			break
		} else {
			break
		}
	}
	if !tilesMoved {
		board.BlockSameMove = UP
	} else {
		board.BlockSameMove = NOOP
	}
}

func (board *Board) shiftDown(nextIndex int) {
	tilesMoved := false
	for {
		prev := nextIndex
		thisCell := board.getSquare(prev)
		if thisCell.Value == 0 {
			return
		}
		nextIndex = nextIndex + 4
		if nextIndex > 15 {
			//already at the lower edge
			break
		}
		next := board.getSquare(nextIndex)
		if next.Value == 0 {
			//sliiiide it.
			board.setSquareValue(nextIndex, thisCell.Value)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
		} else if next.Value == thisCell.Value && !next.AlreadyCombined && !thisCell.AlreadyCombined {
			//combine
			board.combineSquare(nextIndex)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
			break
		} else {
			break
		}
	}
	if !tilesMoved {
		board.BlockSameMove = DOWN
	} else {
		board.BlockSameMove = NOOP
	}
}

func (board *Board) shiftLeft(nextIndex int) {
	tilesMoved := false
	for {
		prev := nextIndex
		thisCell := board.getSquare(prev)
		if thisCell.Value == 0 {
			return
		}
		nextIndex = nextIndex - 1
		if prev == 0 || prev == 4 || prev == 8 || prev == 12 {
			//already at the edge
			break
		}
		next := board.getSquare(nextIndex)
		if next.Value == 0 {
			//sliiiide it.
			board.setSquareValue(nextIndex, thisCell.Value)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
		} else if next.Value == thisCell.Value && !next.AlreadyCombined && !thisCell.AlreadyCombined {
			//combine
			board.combineSquare(nextIndex)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
			break
		} else {
			break
		}
	}
	if !tilesMoved {
		board.BlockSameMove = LEFT
	} else {
		board.BlockSameMove = NOOP
	}
}

func (board *Board) shiftRight(nextIndex int) {
	tilesMoved := false
	for {
		prev := nextIndex
		thisCell := board.getSquare(prev)
		if thisCell.Value == 0 {
			return
		}
		nextIndex = nextIndex + 1
		if prev == 3 || prev == 7 || prev == 11 || prev == 15 {
			//already at the edge
			break
		}
		next := board.getSquare(nextIndex)
		if next.Value == 0 {
			//sliiiide it.
			board.setSquareValue(nextIndex, thisCell.Value)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
		} else if next.Value == thisCell.Value && !next.AlreadyCombined && !thisCell.AlreadyCombined {
			//combine
			board.combineSquare(nextIndex)
			//zero out old spot.
			board.setSquareValue(prev, 0)
			tilesMoved = true
			break
		} else {
			break
		}
	}
	if !tilesMoved {
		board.BlockSameMove = RIGHT
	} else {
		board.BlockSameMove = NOOP
	}
}

func (board *Board) setSquareValue(index, value int) {
	board.getSquare(index).Value = value
}

func (board *Board) combineSquare(index int) {
	next := board.getSquare(index)
	if next.AlreadyCombined {
		panic("Combining already combined square for this generation")
	}
	next.Value = next.Value * 2
	if next.Value == 2048 {
		board.Won = true
	}
	next.AlreadyCombined = true
}

func (board *Board) getSquare(i int) *Square {
	return &board.Grid[i]
}

func (board *Board) Print() {
	fmt.Printf("\033[0;0H") // lets things print in place
	fmt.Println("|-------------------|")
	for i := 0; i <= 12; i += 4 {
		fmt.Printf("|%s|%s|%s|%s|\n", board.getSquare(i), board.getSquare(i+1), board.getSquare(i+2), board.getSquare(i+3))
	}
	fmt.Print("|-------------------|\n\n")

}
