package main

import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type SDLDisplay struct {
	Window *sdl.Window
	Joypad *byte
	IF *ReferenceRegister
	IE *ReferenceRegister
	CPU *CPU
}

var count = 0
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

func NewSDLDisplay(joypad *byte, interrupt *ReferenceRegister) (*SDLDisplay, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	_ = sdl.GLSetSwapInterval(0)

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

	return &SDLDisplay{Window: window, Joypad: joypad, IF: interrupt}, nil
}

func (display SDLDisplay) UpdateScanline(pixels []*Pixel, palette byte, y int) {
	count += 1
	surface, _ := display.Window.GetSurface()
	for x := 0; x < len(pixels); x++ {
		surface.Set(x, y,  getColorFromPixel(pixels[x], palette))
	}
}

func (display *SDLDisplay) handleKeyboardEvent(ev *sdl.KeyboardEvent){
	if ev.State == 0 {
		return
	}
	var joypadIndex int
	switch key := ev.Keysym ; key.Sym {
	case sdl.K_RETURN:
		joypadIndex = SELECT_JOYPAD
	case sdl.K_RIGHT:
		joypadIndex = RIGHT_JOYPAD
	case sdl.K_LEFT:
		joypadIndex = LEFT_JOYPAD
	case sdl.K_UP:
		joypadIndex = UP_JOYPAD
	case sdl.K_DOWN:
		joypadIndex = DOWN_JOYPAD
	case sdl.K_a:
		joypadIndex = A_JOYPAD
	case sdl.K_b:
		joypadIndex = B_JOYPAD
	case sdl.K_SPACE:
		joypadIndex = START_JOYPAD
	default:
		return
	}

	*display.Joypad = SetBit(*display.Joypad, 0, joypadIndex)
	display.IF.SetBit(1, JOYPAD)
}

func (display SDLDisplay) HandleEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			display.handleKeyboardEvent(event.(*sdl.KeyboardEvent))
		}

	}
	return true
}


func (display SDLDisplay) UpdateGraphics() error {
	return display.Window.UpdateSurface()
}
