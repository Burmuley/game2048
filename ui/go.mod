module github.com/Burmuley/game2048/ui

go 1.14

replace github.com/Burmuley/game2048/engine => ../engine

replace github.com/Burmuley/game2048/ui/fyne => ./fyne

replace github.com/Burmuley/game2048/ui/console => ./console

require (
	github.com/Burmuley/game2048/engine v0.0.0-00010101000000-000000000000
	github.com/Burmuley/game2048/ui/console v0.0.0-00010101000000-000000000000
	github.com/Burmuley/game2048/ui/fyne v0.0.0-00010101000000-000000000000
)
