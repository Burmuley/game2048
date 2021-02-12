package engine

import (
	"math/rand"
	"time"
)

type Game2048 struct {
	field    Field
	prevMove int8
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
	var cMove int8
	g.field, cMove = m(g.field)

	if g.prevMove != cMove {
		g.addRandomTwos()
	}

	g.prevMove = cMove
	return g.field
}

func (g *Game2048) addRandomTwos() {
	rcLen := len(g.field) - 1

	for r := range g.field {
		cTime := time.Now()
		rand.Seed(cTime.UnixNano())
		//rndRow := rand.Intn(rcLen)

		for range g.field[r] {
			rndCol := rand.Intn(rcLen)

			if g.field[r][rndCol] == EmptyCell {
				g.field[r][rndCol] = 2
				return
			}
		}
	}

	panic("GAME OVER!")
}
