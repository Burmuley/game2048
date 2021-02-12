package engine

const (
	EmptyCell int = 0
)

type Field [][]int
type Move func(f Field) (Field, int8)

type Engine interface {
	Init(n int) Engine
	Move(m Move) Field
	Field() Field
}

type EngineField interface {
	Get() Field
	Set(Field)
}
