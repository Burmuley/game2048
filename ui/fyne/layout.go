package fyne

import (
	"fyne.io/fyne/v2"
)

type GameWindowLayout struct {
}

func NewGameWindowLayout() *GameWindowLayout {
	return &GameWindowLayout{}
}

func (g *GameWindowLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	// place all elements as vertical BoxLayout does, but scale bottom to fit all space left
	nextPos := fyne.NewPos(0, 0)
	sumHeight := float32(0)
	numObjects := len(objects) - 1

	for i, child := range objects {
		if i < numObjects {
			child.Resize(child.MinSize())
			sumHeight += child.MinSize().Height
			child.Move(nextPos)
			nextPos.Y += sumHeight
			continue
		}

		child.Resize(fyne.NewSize(size.Width, size.Height-sumHeight))
		child.Move(nextPos)
	}
}

func (g *GameWindowLayout) MinSize(objects []fyne.CanvasObject) (minSize fyne.Size) {
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		minSize = minSize.Max(child.MinSize())
	}

	return
}
