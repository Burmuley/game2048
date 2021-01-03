package engine

import (
	"math/rand"
	"time"
)

type Game2048 struct {
	field Field
}

func NewGame2048() *Game2048 {
	return &Game2048{}
}

func (g *Game2048) Field() Field {
	return g.field
}

func (g *Game2048) Init(n int) Engine {
	// make a field of size n
	g.field = make([][]int, n, n)
	for i := range g.field {
		g.field[i] = make([]int, n, n)
		for k := range g.field[i] {
			g.field[i][k] = EmptyCell
		}
	}

	// fill field randomly
	for k := 0; k < n/2; k++ {
		cTime := time.Now()
		rand.Seed(cTime.UnixNano())
		g.field[rand.Intn(n)][rand.Intn(n)] = 2
	}

	return g
}

func (g *Game2048) Move(m Move) Field {
	g.field = m(g.field)
	return g.field
}
