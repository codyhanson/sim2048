package game

import "testing"

func TestNew(t *testing.T) {
	board := New()
	if len(board.Grid) != 16 {
		t.Error("Grid length != 16")
	}
	numberOfTwos := 0
	for _, v := range board.Grid {
		if v.Value == 2 {
			numberOfTwos++
			continue
		}
		if v.Value != 0 {
			t.Error("Board seeded with something besides a 2 or a 0")
		}
	}
	if numberOfTwos != 2 {
		t.Error("Board seeded not exactly 2 2's")
	}
}

func TestAdvanceGrid(t *testing.T) {
	board := New()
	board.Grid[1].AlreadyCombined = true
	board.AdvanceGrid()
	if board.Grid[1].AlreadyCombined {
		t.Error("AlreadyCombined not set to false.")
	}
}

func TestGenNewSquares(t *testing.T) {

}

func TestShiftUp(t *testing.T) {
	board := &Board{
		Grid:     [16]Square{},
	}

	// Starting state
	//2  4  2  0
	//2  0  4  4
	//0  8  2  0
	//2  4  2  2

	board.Grid[0].Value = 2
	board.Grid[1].Value = 4
	board.Grid[2].Value = 2
	board.Grid[3].Value = 0
	board.Grid[4].Value = 2
	board.Grid[5].Value = 0
	board.Grid[6].Value = 4
	board.Grid[7].Value = 4
	board.Grid[8].Value = 0
	board.Grid[9].Value = 8
	board.Grid[10].Value = 2
	board.Grid[11].Value = 0
	board.Grid[12].Value = 2
	board.Grid[13].Value = 4
	board.Grid[14].Value = 2
	board.Grid[15].Value = 2
	board.Shift(UP)
	// Becomes
	//4  4  2  4
	//2  8  4  2
	//0  4  4  0
	//0  0  0  0
	assertValue(board, 0, 4, t)
	assertValue(board, 1, 4, t)
	assertValue(board, 2, 2, t)
	assertValue(board, 3, 4, t)
	assertValue(board, 4, 2, t)
	assertValue(board, 5, 8, t)
	assertValue(board, 6, 4, t)
	assertValue(board, 7, 2, t)
	assertValue(board, 8, 0, t)
	assertValue(board, 9, 4, t)
	assertValue(board, 10, 4, t)
	assertValue(board, 11, 0, t)
	assertValue(board, 12, 0, t)
	assertValue(board, 13, 0, t)
	assertValue(board, 14, 0, t)
	assertValue(board, 15, 0, t)
}


func TestShiftRight(t *testing.T) {
	board := &Board{
		Grid:     [16]Square{},
	}

	// Starting state
	//2  4  2  0
	//2  0  4  4
	//0  8  2  0
	//2  4  2  2

	board.Grid[0].Value = 2
	board.Grid[1].Value = 4
	board.Grid[2].Value = 2
	board.Grid[3].Value = 0
	board.Grid[4].Value = 2
	board.Grid[5].Value = 0
	board.Grid[6].Value = 4
	board.Grid[7].Value = 4
	board.Grid[8].Value = 0
	board.Grid[9].Value = 8
	board.Grid[10].Value = 2
	board.Grid[11].Value = 0
	board.Grid[12].Value = 2
	board.Grid[13].Value = 4
	board.Grid[14].Value = 2
	board.Grid[15].Value = 2
	board.Shift(RIGHT)
	// Becomes
	//0 2 4 2
	//0 0 2 8
	//0 0 8 2
	//0 2 4 4
	assertValue(board, 0, 0, t)
	assertValue(board, 1, 2, t)
	assertValue(board, 2, 4, t)
	assertValue(board, 3, 2, t)
	assertValue(board, 4, 0, t)
	assertValue(board, 5, 0, t)
	assertValue(board, 6, 2, t)
	assertValue(board, 7, 8, t)
	assertValue(board, 8, 0, t)
	assertValue(board, 9, 0, t)
	assertValue(board, 10, 8, t)
	assertValue(board, 11, 2, t)
	assertValue(board, 12, 0, t)
	assertValue(board, 13, 2, t)
	assertValue(board, 14, 4, t)
	assertValue(board, 15, 4, t)
}


func TestShiftLeft(t *testing.T) {
	board := &Board{
		Grid:     [16]Square{},
	}

	// Starting state
	//2  4  2  0
	//2  0  4  4
	//0  8  2  0
	//2  4  2  2

	board.Grid[0].Value = 2
	board.Grid[1].Value = 4
	board.Grid[2].Value = 2
	board.Grid[3].Value = 0
	board.Grid[4].Value = 2
	board.Grid[5].Value = 0
	board.Grid[6].Value = 4
	board.Grid[7].Value = 4
	board.Grid[8].Value = 0
	board.Grid[9].Value = 8
	board.Grid[10].Value = 2
	board.Grid[11].Value = 0
	board.Grid[12].Value = 2
	board.Grid[13].Value = 4
	board.Grid[14].Value = 2
	board.Grid[15].Value = 2
	board.Shift(LEFT)
	// Becomes
	//2 4 2 0
	//2 8 0 0
	//8 2 0 0
	//2 4 4 0
	assertValue(board, 0, 2, t)
	assertValue(board, 1, 4, t)
	assertValue(board, 2, 2, t)
	assertValue(board, 3, 0, t)
	assertValue(board, 4, 2, t)
	assertValue(board, 5, 8, t)
	assertValue(board, 6, 0, t)
	assertValue(board, 7, 0, t)
	assertValue(board, 8, 8, t)
	assertValue(board, 9, 2, t)
	assertValue(board, 10, 0, t)
	assertValue(board, 11, 0, t)
	assertValue(board, 12, 2, t)
	assertValue(board, 13, 4, t)
	assertValue(board, 14, 4, t)
	assertValue(board, 15, 0, t)
}


func TestShiftDown(t *testing.T) {
	board := &Board{
		Grid:     [16]Square{},
	}

	// Starting state
	//2  4  2  4
	//2  0  4  0
	//0  8  2  0
	//2  4  2  2

	board.Grid[0].Value = 2
	board.Grid[1].Value = 4
	board.Grid[2].Value = 2
	board.Grid[3].Value = 4
	board.Grid[4].Value = 2
	board.Grid[5].Value = 0
	board.Grid[6].Value = 4
	board.Grid[7].Value = 0
	board.Grid[8].Value = 0
	board.Grid[9].Value = 8
	board.Grid[10].Value = 2
	board.Grid[11].Value = 0
	board.Grid[12].Value = 2
	board.Grid[13].Value = 4
	board.Grid[14].Value = 2
	board.Grid[15].Value = 2
	board.Print()
	board.Shift(DOWN)
	// Becomes
	// 0 0 0 0
	// 0 4 2 0
	// 2 8 4 4
	// 4 4 4 2

	board.Print()
	assertValue(board, 0, 0, t)
	assertValue(board, 1, 0, t)
	assertValue(board, 2, 0, t)
	assertValue(board, 3, 0, t)
	assertValue(board, 4, 0, t)
	assertValue(board, 5, 4, t)
	assertValue(board, 6, 2, t)
	assertValue(board, 7, 0, t)
	assertValue(board, 8, 2, t)
	assertValue(board, 9, 8, t)
	assertValue(board, 10, 4, t)
	assertValue(board, 11, 4, t)
	assertValue(board, 12, 4, t)
	assertValue(board, 13, 4, t)
	assertValue(board, 14, 4, t)
	assertValue(board, 15, 2, t)
}

func assertValue(board *Board, index int, expected int, t *testing.T) {
	if board.Grid[index].Value != expected {
		t.Errorf("Expected index %d to have value %d but got %d", index, expected, board.Grid[index].Value)
	}

}
