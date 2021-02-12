package fyne

import "image/color"

var (
	numberColors = map[int]color.Color{
		0:    color.RGBA{0, 0, 0, 0},
		2:    color.RGBA{136, 180, 222, 255},
		4:    color.RGBA{88, 148, 207, 255},
		8:    color.RGBA{61, 104, 145, 255},
		16:   color.RGBA{237, 190, 124, 255},
		32:   color.RGBA{231, 163, 75, 255},
		64:   color.RGBA{196, 138, 62, 255},
		128:  color.RGBA{213, 118, 89, 255},
		256:  color.RGBA{207, 96, 63, 255},
		512:  color.RGBA{0, 0, 0, 0},
		1024: color.RGBA{0, 0, 0, 0},
		2048: color.RGBA{0, 0, 0, 0},
	}
)
