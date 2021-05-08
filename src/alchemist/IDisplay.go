package main

import "image/color"

type IDisplay interface {
	UpdateScanline(pixels []*Pixel, palette byte, ly int)
	UpdateScanlinePixels(colors []color.RGBA, ly int)
	HandleEvent() bool
	UpdateGraphics() error
}
