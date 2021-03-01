package engine

const (
	EmptyCell int = 0
	Newcell   int = 2
	WinNumber int = 2048
)

const (
	FAL_NO FieldAlingment = iota
	FAL_LEFT
	FAL_RIGHT
	FAL_UP
	FAL_DOWN
)

type FieldAlingment int8
type Move func(f Field) MoveResult

type Field interface {
	Get() [][]int
	Set([][]int)
	HasFreeCells() bool
}

type Engine interface {
	Move(m Move) (MoveResult, error)
	Field() [][]int
	AddScore(inc int)
	Score() int
}

type MoveResult interface {
	Aligned() FieldAlingment
	Changed() bool
	Score() int
}
