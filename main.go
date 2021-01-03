package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Burmuley/game2048/engine"
	"github.com/Burmuley/game2048/ui"
)

func isStrInList(s string, l []string) bool {
	for _, v := range l {
		if s == v {
			return true
		}
	}

	return false
}

func main() {
	var cmdUIName string
	var eng engine.Engine

	flag.StringVar(&cmdUIName, "ui", "gui", "choose UI: 'console' or 'gui'")
	flag.Parse()

	uiFabric := ui.NewFabric()

	if !isStrInList(cmdUIName, uiFabric.List()) {
		fmt.Printf("Unknown UI defined: %s\n", cmdUIName)
		os.Exit(1)
	}

	eng = engine.NewGame2048()
	eng = eng.Init(4)

	ui := uiFabric.Get(ui.Gui)
	ui.SetGame(eng)
	ui.Run()
}
