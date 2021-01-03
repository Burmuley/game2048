package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Burmuley/game2048/engine"
	"github.com/Burmuley/game2048/ui"
)

const (
	defaultUI = ui.Gui
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

	uiFabric := ui.NewFabric()

	flag.StringVar(
		&cmdUIName,
		"ui",
		defaultUI,
		fmt.Sprintf(
			"choose UI: %s",
			strings.Join(uiFabric.List(), " | ")))
	flag.Parse()

	if !isStrInList(cmdUIName, uiFabric.List()) {
		fmt.Printf("Unknown UI defined: %s\n", cmdUIName)
		os.Exit(1)
	}

	eng = engine.NewGame2048()
	eng = eng.Init(4)

	appUI := uiFabric.Get(cmdUIName)
	appUI.SetGame(eng)
	appUI.Run()
}
