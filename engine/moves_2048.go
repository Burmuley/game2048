package engine

type Move func(f Field) Field

var (
	M_RIGHT2048 Move = Move2048Right
	M_LEFT2048  Move = Move2048Left
	M_UP2048    Move = Move2048Up
	M_DOWN2048  Move = Move2048Down
)

func Move2048Right(f Field) Field {
	last := len(f) - 1
	sumCounter := 0
	maxSum := 2

	for i := range f {
		for b := 0; b < last; b++ {
			for k := last; k >= 1; k-- {
				if f[i][k] == EmptyCell {
					f[i][k], f[i][k-1] = f[i][k-1], f[i][k]
					continue
				}

				if f[i][k] == f[i][k-1] && sumCounter < maxSum {
					f[i][k], f[i][k-1] = f[i][k]*2, EmptyCell
					sumCounter++
					continue
				}
			}
		}
	}

	return f
}

func Move2048Left(f Field) Field {
	last := len(f) - 1
	sumCounter := 0
	maxSum := 2

	for i := range f {
		for b := 0; b < last; b++ {
			for k := 0; k < last; k++ {
				if f[i][k] == EmptyCell {
					f[i][k], f[i][k+1] = f[i][k+1], f[i][k]
					continue
				}

				if f[i][k] == f[i][k+1] && sumCounter < maxSum {
					f[i][k], f[i][k+1] = f[i][k]*2, EmptyCell
					sumCounter++
					continue
				}
			}
		}
	}

	return f
}

func Move2048Up(f Field) Field {
	last := len(f) - 1
	sumCounter := 0
	maxSum := 2

	for i := range f {
		for b := 0; b < last; b++ {
			for k := 0; k < last; k++ {
				if f[k][i] == EmptyCell {
					f[k][i], f[k+1][i] = f[k+1][i], f[k][i]
					continue
				}

				if f[k][i] == f[k+1][i] && sumCounter < maxSum {
					f[k][i], f[k+1][i] = f[k][i]*2, EmptyCell
					sumCounter++
					continue
				}
			}
		}
	}

	return f
}

func Move2048Down(f Field) Field {
	last := len(f) - 1
	sumCounter := 0
	maxSum := 2

	for i := range f {
		for b := 0; b < last; b++ {
			for k := last; k >= 1; k-- {
				if f[k][i] == EmptyCell {
					f[k][i], f[k-1][i] = f[k-1][i], f[k][i]
					continue
				}

				if f[k][i] == f[k-1][i] && sumCounter < maxSum {
					f[k][i], f[k-1][i] = f[k][i]*2, EmptyCell
					sumCounter++
					continue
				}
			}
		}
	}

	return f
}
