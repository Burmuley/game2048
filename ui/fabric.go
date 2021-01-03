package ui

import (
	"github.com/Burmuley/game2048/ui/console"
	"github.com/Burmuley/game2048/ui/fyne"
)

const (
	Console string = "console"
	Gui     string = "gui"
)

func getUI(n string) UI {
	switch n {
	case Console:
		return console.New()
	case Gui:
		return fyne.New()
	default:
		return nil
	}
}

type Fabric struct {
	uis    []string
	uiFunc func(string) UI
}

func NewFabric() *Fabric {
	f := &Fabric{
		uis: []string{
			"gui",
			"console",
		},
		uiFunc: getUI,
	}
	return f
}

func (f *Fabric) Get(ui string) UI {
	return f.uiFunc(ui)
}

func (f *Fabric) List() []string {
	return f.uis
}
