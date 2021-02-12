package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Burmuley/game2048/engine"
)

type GameScore struct {
	widget.Label
	game          engine.Engine
	gameContainer *fyne.Container
	ui            *FyneUI
	preKey        fyne.KeyName
}

func (g *GameScore) FocusGained() {
	return
}

func (g *GameScore) FocusLost() {
	return
}

func (g *GameScore) TypedRune(r rune) {
	return
}

func (g *GameScore) TypedKey(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeyRight:
		g.game.Move(engine.M_RIGHT2048)
	case fyne.KeyLeft:
		g.game.Move(engine.M_LEFT2048)
	case fyne.KeyUp:
		g.game.Move(engine.M_UP2048)
	case fyne.KeyDown:
		g.game.Move(engine.M_DOWN2048)
	}

	g.gameContainer.Refresh()
}

func NewGameScore(text string, game engine.Engine, gameContainer *fyne.Container, ui *FyneUI) *GameScore {
	score := &GameScore{
		Label:         widget.Label{Text: text},
		game:          game,
		gameContainer: gameContainer,
		ui:            ui,
	}
	score.ExtendBaseWidget(score)
	return score
}
