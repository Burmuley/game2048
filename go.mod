module github.com/Burmuley/game2048

go 1.14

replace github.com/Burmuley/game2048/engine => ./engine

replace github.com/Burmuley/game2048/ui => ./ui

replace github.com/Burmuley/game2048/ui/fyne => ./ui/fyne

replace github.com/Burmuley/game2048/ui/console => ./ui/console

require (
	github.com/Burmuley/game2048/engine v0.0.0-00010101000000-000000000000
	github.com/Burmuley/game2048/ui v0.0.0-00010101000000-000000000000
)
