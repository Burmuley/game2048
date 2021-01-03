package ui

import "github.com/Burmuley/game2048/engine"

type UI interface {
	SetGame(engine engine.Engine)
	Name() string
	Run()
}
