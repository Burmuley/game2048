package engine

var (
	M_RIGHT2048 Move = Move2048Right
	M_LEFT2048  Move = Move2048Left
	M_UP2048    Move = Move2048Up
	M_DOWN2048  Move = Move2048Down
)

func Move2048Right(field Field) MoveResult {
	f := field.Get()
	startF := copyField(f)
	last := len(f) - 1
	score := 0

	// shift all cells to the right
	shiftRight(f)

	// sum neighbor cells
	for r := range f {
		for c := last; c >= 1; c-- {
			if f[r][c] == f[r][c-1] {
				newVal := f[r][c] * 2
				f[r][c], f[r][c-1] = newVal, EmptyCell
				score += newVal
			}
		}
	}

	// shift all cells to the right again
	// to remove gaps after summing
	shiftRight(f)
	field.Set(f)
	return NewGame2048Result(FAL_RIGHT, !cmpFields(startF, f), score)
}

func Move2048Left(field Field) MoveResult {
	f := field.Get()
	startF := copyField(f)
	last := len(f) - 1
	score := 0

	// shift all cells to the left
	shiftLeft(f)

	// sum neighbor cells
	for r := range f {
		for c := 0; c < last; c++ {
			if f[r][c] == f[r][c+1] {
				newVal := f[r][c] * 2
				f[r][c], f[r][c+1] = newVal, EmptyCell
				score += newVal
			}
		}
	}

	// shift all cells to the left again
	// to remove gaps after summing
	shiftLeft(f)
	field.Set(f)
	return NewGame2048Result(FAL_LEFT, !cmpFields(startF, f), score)
}

func Move2048Up(field Field) MoveResult {
	f := field.Get()
	startF := copyField(f)
	last := len(f) - 1
	score := 0

	// shift all cells up
	shiftUp(f)

	// sum neighbor cells
	for r := 0; r < last; r++ {
		for c := range f[r] {
			if f[r][c] == f[r+1][c] {
				newVal := f[r][c] * 2
				f[r][c], f[r+1][c] = newVal, EmptyCell
				score += newVal
			}
		}
	}

	// shift all cells up again
	// to remove gaps after summing
	shiftUp(f)
	field.Set(f)
	return NewGame2048Result(FAL_UP, !cmpFields(startF, f), score)
}

func Move2048Down(field Field) MoveResult {
	f := field.Get()
	startF := copyField(f)
	last := len(f) - 1
	score := 0

	// shift all cells down
	shiftDown(f)

	// sum neighbor cells
	for c := range f {
		for r := last; r > 0; r-- {
			if f[r][c] == f[r-1][c] {
				newVal := f[r][c] * 2
				f[r][c], f[r-1][c] = newVal, EmptyCell
				score += newVal
			}
		}
	}

	// shift all cells down again
	// to remove gaps after summing
	shiftDown(f)
	field.Set(f)
	return NewGame2048Result(FAL_DOWN, !cmpFields(startF, f), score)
}

// Helper functions
func shiftRight(field [][]int) {
	for range field {
		for _, r := range field {
			for i := len(r) - 1; i > 0; i-- {
				if r[i] == EmptyCell && r[i-1] == EmptyCell {
					continue
				}
				if r[i] == EmptyCell {
					r[i], r[i-1] = r[i-1], r[i]
				}
			}
		}
	}
}

func shiftLeft(field [][]int) {
	for range field {
		for _, r := range field {
			for i := 0; i < len(r)-1; i++ {
				if r[i] == EmptyCell && r[i+1] == EmptyCell {
					continue
				}
				if r[i] == EmptyCell {
					r[i], r[i+1] = r[i+1], r[i]
				}
			}
		}
	}
}

func shiftUp(field [][]int) {
	last := len(field) - 1

	for range field {
		for r := 0; r < last; r++ {
			for c := range field[r] {
				if field[r][c] == EmptyCell && field[r+1][c] == EmptyCell {
					continue
				}
				if field[r][c] == EmptyCell {
					field[r][c], field[r+1][c] = field[r+1][c], field[r][c]
				}
			}
		}
	}
}

func shiftDown(field [][]int) {
	last := len(field) - 1

	for range field {
		for r := last; r > 0; r-- {
			for c := range field {
				if field[r][c] == EmptyCell && field[r-1][c] == EmptyCell {
					continue
				}
				if field[r][c] == EmptyCell {
					field[r][c], field[r-1][c] = field[r-1][c], field[r][c]
				}
			}
		}
	}
}

// compare two fields
// return true if they're equal
func cmpFields(before, after [][]int) bool {
	for r := range before {
		for c := range before[r] {
			if before[r][c] != after[r][c] {
				return false
			}
		}
	}

	return true
}

// make and return a copy of the game field
func copyField(field [][]int) [][]int {
	cp := make([][]int, len(field), len(field))

	for r := range cp {
		cp[r] = make([]int, len(field), len(field))
	}
	for r := range field {
		copy(cp[r], field[r])
	}

	return cp
}
