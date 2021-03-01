package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/Burmuley/game2048/engine"
)

const (
	cellSize int = 100
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
	uiApp := app.New()
	uiWindow := uiApp.NewWindow("Game 2048")
	fSize := len(f.game.Field())

	gameLayout := layout.NewGridLayout(len(f.game.Field()))
	gameContainer := container.New(gameLayout)

	scoreLabel := NewGameScore("Score: 0", f.game, gameContainer, f, uiWindow)
	scoreLabel.Resize(scoreLabel.MinSize())
	scoreLabel.Alignment = fyne.TextAlignCenter

	gridSize := fyne.NewSize(
		float32(fSize*cellSize),
		float32(fSize*cellSize),
	)
	wSize := fyne.Size{
		Width:  gridSize.Width,
		Height: gridSize.Height + scoreLabel.MinSize().Height,
	}

	wLayout := NewGameWindowLayout()
	wContainer := container.New(wLayout)
	wContainer.Add(scoreLabel)
	wContainer.Add(gameContainer)
	wContainer.Resize(wSize)
	f.fillContainer(gameContainer)

	uiWindow.Resize(wSize)
	uiWindow.SetContent(wContainer)
	uiWindow.Canvas().Focus(scoreLabel)
	uiWindow.ShowAndRun()
}

func New() *FyneUI {
	return &FyneUI{
		name: "Fyne",
	}
}

func (f *FyneUI) fillContainer(cont *fyne.Container) {
	cont.Objects = make([]fyne.CanvasObject, 0, 0)
	for r := range f.game.Field() {
		for c := range f.game.Field()[r] {
			t := NewGameCell(f.game, [2]int{r, c})
			cont.Add(t)
		}
	}
}
