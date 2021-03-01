package engine

import (
	"errors"
	"math/rand"
	"time"
)

type Game2048Result struct {
	aligned FieldAlingment
	changed bool
	score   int
}

func NewGame2048Result(aligned FieldAlingment, changed bool, score int) *Game2048Result {
	return &Game2048Result{aligned: aligned, changed: changed, score: score}
}

func (g Game2048Result) Aligned() FieldAlingment {
	return g.aligned
}

func (g Game2048Result) Changed() bool {
	return g.changed
}

func (g Game2048Result) Score() int {
	return g.score
}

type Game2048 struct {
	field Field
	score int
}

func NewGame2048(size int) *Game2048 {
	field := NewGame2048Field(size)
	return &Game2048{
		field: field,
		score: 0,
	}
}

func NewGame2048WithField(field Field) *Game2048 {
	return &Game2048{
		field: field,
		score: 0,
	}
}

func (g *Game2048) Field() [][]int {
	return g.field.Get()
}

func (g *Game2048) Move(move Move) (MoveResult, error) {
	var cMove MoveResult
	cMove = move(g.field)

	if !g.field.HasFreeCells() {
		return cMove, errors.New("GAME OVER!")
	}

	if g.checkWin() {
		return cMove, errors.New("IT'S A WIN!")
	}

	if cMove.Changed() {
		g.addRandomTwos(cMove.Aligned())
	}

	return cMove, nil
}

func (g *Game2048) AddScore(inc int) {
	g.score += inc
}

func (g *Game2048) Score() int {
	return g.score
}

func (g *Game2048) addRandomTwos(aligned FieldAlingment) {
	field := g.field.Get()
	fSize := (len(field) - 1) * 10
	rndSize := len(field) * 100

	for i := 0; i < fSize; i++ {
		cTime := time.Now()
		rand.Seed(cTime.UnixNano())
		rndRow := int(rand.Intn(rndSize) / 100)
		rndCol := int(rand.Intn(rndSize) / 100)

		if field[rndRow][rndCol] == EmptyCell {
			field[rndRow][rndCol] = Newcell
			break
		}
	}

	g.field.Set(field)
}

func (g *Game2048) checkWin() bool {
	for _, row := range g.field.Get() {
		for cell := range row {
			if cell == WinNumber {
				return true
			}
		}
	}

	return false
}

// Field
type Game2048Field struct {
	aligned FieldAlingment
	field   [][]int
	changed bool
}

func (g *Game2048Field) Get() [][]int {
	return g.field
}

func (g *Game2048Field) Set(f [][]int) {
	g.field = f
}

func (g *Game2048Field) HasFreeCells() bool {
	for r := range g.field {
		for _, c := range g.field[r] {
			if c == EmptyCell {
				return true
			}
		}
	}

	return false
}

func NewGame2048Field(size int) Field {
	fieldInstance := Game2048Field{
		aligned: FAL_NO,
		changed: false,
	}

	field := make([][]int, size, size)

	for i := range field {
		field[i] = make([]int, size, size)
		for k := range field[i] {
			field[i][k] = EmptyCell
		}
	}

	// fill field randomly
	for k := 0; k < size/2; k++ {
		cTime := time.Now()
		rand.Seed(cTime.UnixNano())
		field[rand.Intn(size)][rand.Intn(size)] = Newcell
	}

	fieldInstance.field = field

	return &fieldInstance
}
