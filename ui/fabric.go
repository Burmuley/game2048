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
	uis map[string]func(string) UI
}

func NewFabric() *Fabric {
	f := &Fabric{
		uis: make(map[string]func(string) UI, 0),
	}

	f.Add(Console, getUI)
	f.Add(Gui, getUI)

	return f
}

func (f *Fabric) Get(ui string) UI {
	return f.uis[ui](ui)
}

func (f *Fabric) Add(n string, fn func(string) UI) {
	f.uis[n] = fn
}

func (f *Fabric) List() []string {
	l := make([]string, 0)
	for i := range f.uis {
		l = append(l, i)
	}

	return l
}
