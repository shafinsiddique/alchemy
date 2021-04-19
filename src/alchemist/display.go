package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type SDLDisplay struct {
	Window *sdl.Window
}

func getColor(first byte, second byte) color.RGBA {
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

func getColorFromPixel(pixel *Pixel, palette byte) color.RGBA {
	switch current := pixel.Get(); current {
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

func clearBackground(window *sdl.Window) error {
	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	for x := 0; x < WINDOW_WIDTH; x++ {
		for y := 0; y < WINDOW_HEIGHT; y++ {
			surface.Set(x, y, COLOR_MAP[0])
		}
	}
	err = window.UpdateSurface()
	return err
}

func NewSDLDisplay() (*SDLDisplay, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow(TITLE,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	err = clearBackground(window)

	if err != nil {
		return nil, err
	}

	return &SDLDisplay{Window: window}, nil
}

func (display SDLDisplay) UpdateScanline(pixels []*Pixel, palette byte, y int) {
	surface, _ := display.Window.GetSurface()
	for x := 0; x < len(pixels); x++ {
		surface.Set(x, y, getColorFromPixel(pixels[x], palette))
	}
	_ = display.UpdateGraphics()
}

func (display SDLDisplay) HandleEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false
		}
	}
	return true
}


func (display SDLDisplay) UpdateGraphics() error {
	return display.Window.UpdateSurface()
}
