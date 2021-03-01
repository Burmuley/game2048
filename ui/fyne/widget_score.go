package fyne

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Burmuley/game2048/engine"
)

type GameScore struct {
	widget.Label
	game          engine.Engine
	gameContainer *fyne.Container
	ui            *FyneUI
	uiWindow      fyne.Window
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
	var err error
	var res engine.MoveResult

	switch event.Name {
	case fyne.KeyRight:
		res, err = g.game.Move(engine.M_RIGHT2048)
	case fyne.KeyLeft:
		res, err = g.game.Move(engine.M_LEFT2048)
	case fyne.KeyUp:
		res, err = g.game.Move(engine.M_UP2048)
	case fyne.KeyDown:
		res, err = g.game.Move(engine.M_DOWN2048)
	}

	g.game.AddScore(res.Score())

	if err != nil {
		dlg := dialog.NewCustom("Message:", "Another shot!", canvas.NewText(fmt.Sprintf("%s", err), color.White), g.uiWindow)
		dlg.SetOnClosed(func() {
			g.game = engine.NewGame2048(4)
			g.ui.game = g.game
			g.ui.fillContainer(g.gameContainer)
			g.Text = fmt.Sprintf("Score: %d", g.game.Score())
		})
		dlg.Show()
	}

	g.gameContainer.Refresh()
	g.Text = fmt.Sprintf("Score: %d", g.game.Score())
	g.Refresh()
}

func NewGameScore(text string, game engine.Engine, gameContainer *fyne.Container, ui *FyneUI, win fyne.Window) *GameScore {
	score := &GameScore{
		Label:         widget.Label{Text: text},
		game:          game,
		gameContainer: gameContainer,
		ui:            ui,
		uiWindow:      win,
	}
	score.ExtendBaseWidget(score)
	return score
}
