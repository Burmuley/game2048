package fyne

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/Burmuley/game2048/engine"

	"fyne.io/fyne/v2/widget"
)

type GameCell struct {
	widget.BaseWidget
	game       engine.Engine
	cellCoords [2]int
}

func NewGameCell(game engine.Engine, coords [2]int) *GameCell {
	return &GameCell{game: game, cellCoords: coords}
}

func (g *GameCell) CreateRenderer() fyne.WidgetRenderer {
	g.ExtendBaseWidget(g)
	text := canvas.NewText(g.getText(), color.White)
	text.TextStyle.Bold = true
	background := canvas.NewRectangle(g.getBgColor())
	background.StrokeColor = color.White
	background.StrokeWidth = 0.3

	r := &gameCellRenderer{
		background: background,
		cell:       g,
		label:      text,
		layout:     layout.NewHBoxLayout(),
	}

	return r
}

func (g *GameCell) getNum() int {
	row, col := g.cellCoords[0], g.cellCoords[1]
	num := g.game.Field()[row][col]

	return num
}

func (g *GameCell) getText() string {
	num := g.getNum()
	text := ""

	if num > 0 {
		text = strconv.Itoa(num)
	}

	return text
}

func (g *GameCell) getBgColor() color.Color {
	num := g.getNum()
	color := numberColors[num]
	return color
}

// renderer
type gameCellRenderer struct {
	cell       *GameCell
	background *canvas.Rectangle
	label      *canvas.Text
	layout     fyne.Layout
}

func (r *gameCellRenderer) Destroy() {
	return
}

func (r *gameCellRenderer) Layout(size fyne.Size) {
	var labelPos fyne.Position

	r.background.Resize(size)
	labelSize := r.label.MinSize()
	labelPos.Y = (size.Height - labelSize.Height) / 2
	labelPos.X = (size.Width - labelSize.Width) / 2
	r.label.Move(labelPos)
	r.label.Resize(labelSize)
}

func (r *gameCellRenderer) MinSize() (size fyne.Size) {
	labelSize := r.label.MinSize()
	size.Width = labelSize.Width
	size.Width += theme.Padding()
	size.Height = labelSize.Height
	return
}

func (r *gameCellRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.label}
}

func (r *gameCellRenderer) Refresh() {
	r.label.Text = r.cell.getText()
	r.background.FillColor = r.cell.getBgColor()
	r.background.Refresh()
	r.Layout(r.cell.Size())
	canvas.Refresh(r.cell)
}
