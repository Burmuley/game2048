package fyne

import (
	"github.com/Burmuley/game2048/engine"
)

type FyneUI struct {
	game engine.Engine
	name string
}

func (f *FyneUI) SetGame(e engine.Engine) {
	f.game = e
}

func (f *FyneUI) Name() string {
	return f.name
}

func (f *FyneUI) Run() {
	panic("implement me")
}

func New() *FyneUI {
	return &FyneUI{
		name: "Fyne",
	}
}
