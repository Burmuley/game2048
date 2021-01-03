package engine

const (
	EmptyCell int = 0
)

type Field [][]int

type Engine interface {
	Init(n int) Engine
	Move(m Move) Field
	Field() Field
}
