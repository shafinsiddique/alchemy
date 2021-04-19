package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type Display struct {
	Window sdl.Window
}


func getColor(first byte, second byte) color.RGBA{
	switch first {
	case 1:
		switch second {
		case 1:
			return COLOR_MAP[3]
		default:
			return COLOR_MAP[2]
		}
	default:
		return COLOR_MAP[second]
	}
}


func getColorFromPixel(pixel *Pixel, palette byte) color.RGBA{
	switch current := pixel.Get() ; current {
	case 0:
		return getColor(GetBit(palette, 0), GetBit(palette, 1))
	case 1:
		return getColor(GetBit(palette, 2), GetBit(palette, 3))
	case 2:
		return getColor(GetBit(palette, 4), GetBit(palette, 5))
	default:
		return getColor(GetBit(palette, 6), GetBit(palette, 7))
	}

}

func (display *Display) UpdateScanline(pixels []*Pixel, palette byte, y int) {
	surface, _ := display.Window.GetSurface()
	for x := 0; x < len(pixels) ; x++ {
		surface.Set(x, y, getColorFromPixel(pixels[x], palette))
	}
}
