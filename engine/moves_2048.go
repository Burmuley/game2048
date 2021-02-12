package engine

var (
	M_RIGHT2048 Move = Move2048Right
	M_LEFT2048  Move = Move2048Left
	M_UP2048    Move = Move2048Up
	M_DOWN2048  Move = Move2048Down
)

const (
	MOVE_LEFT int8 = iota
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
)

func Move2048Right(f Field) (Field, int8) {
	last := len(f) - 1

	// shift all cells to the right
	for r := range f {
		f[r] = shiftRight(f[r])
	}

	// sum neighbor cells
	for r := range f {
		for c := last; c >= 1; c-- {
			if f[r][c] == f[r][c-1] {
				f[r][c], f[r][c-1] = f[r][c]*2, EmptyCell
			}
		}
	}

	// shift all cells to the right again
	// to remove gaps after summing
	for r := range f {
		f[r] = shiftRight(f[r])
	}

	return f, MOVE_RIGHT
}

func Move2048Left(f Field) (Field, int8) {
	last := len(f) - 1

	// shift all cells to the left
	for r := range f {
		f[r] = shiftLeft(f[r])
	}

	// sum neighbor cells
	for r := range f {
		for c := 0; c < last; c++ {
			if f[r][c] == f[r][c+1] {
				f[r][c], f[r][c+1] = f[r][c]*2, EmptyCell
			}
		}
	}

	// shift all cells to the left again
	// to remove gaps after summing
	for r := range f {
		f[r] = shiftLeft(f[r])
	}

	return f, MOVE_LEFT
}

func Move2048Up(f Field) (Field, int8) {
	last := len(f) - 1

	// shift all cells up
	for range f {
		f = shiftUp(f)
	}

	// sum neighbor cells
	for r := 0; r < last; r++ {
		for c := range f[r] {
			if f[r][c] == f[r+1][c] {
				f[r][c], f[r+1][c] = f[r][c]*2, EmptyCell
			}
		}
	}

	// shift all cells up again
	// to remove gaps after summing
	for range f {
		f = shiftUp(f)
	}

	return f, MOVE_UP
}

func Move2048Down(f Field) (Field, int8) {
	last := len(f) - 1

	// shift all cells down
	for range f {
		f = shiftDown(f)
	}

	// sum neighbor cells
	for c := range f {
		for r := last; r > 0; r-- {
			if f[r][c] == f[r-1][c] {
				f[r][c], f[r-1][c] = f[r][c]*2, EmptyCell
			}
		}
	}

	// shift all cells down again
	// to remove gaps after summing
	for range f {
		f = shiftDown(f)
	}

	return f, MOVE_DOWN
}

// Helper functions
func shiftRight(row []int) []int {
	for range row {
		for i := len(row) - 1; i > 0; i-- {
			if row[i] == EmptyCell && row[i-1] != EmptyCell {
				row[i], row[i-1] = row[i-1], row[i]
			}
		}
	}

	return row
}

func shiftLeft(row []int) []int {
	for range row {
		for i := 0; i < len(row)-1; i++ {
			if row[i] == EmptyCell && row[i+1] != EmptyCell {
				row[i], row[i+1] = row[i+1], row[i]
			}
		}
	}

	return row
}

func shiftUp(field Field) Field {
	last := len(field) - 1

	for r := 0; r < last; r++ {
		for c := range field[r] {
			if field[r][c] == EmptyCell && field[r+1][c] != EmptyCell {
				field[r][c], field[r+1][c] = field[r+1][c], field[r][c]
			}
		}
	}

	return field
}

func shiftDown(field Field) Field {
	last := len(field) - 1

	for r := last; r > 0; r-- {
		for c := range field {
			if field[r][c] == EmptyCell && field[r-1][c] != EmptyCell {
				field[r][c], field[r-1][c] = field[r-1][c], field[r][c]
			}
		}
	}
	return field
}
