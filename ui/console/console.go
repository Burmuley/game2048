package console

import (
	"fmt"
	"os"

	"github.com/Burmuley/game2048/engine"
)

type Console struct {
	game engine.Engine
	name string
}

func (c *Console) SetGame(g engine.Engine) {
	c.game = g
}

func (c *Console) Name() string {
	return c.name
}

func (c *Console) Run() {
	fmt.Println("NOT IMPLEMENTED YET")
	os.Exit(1)
}

func New() *Console {
	return &Console{
		name: "Console",
	}
}
