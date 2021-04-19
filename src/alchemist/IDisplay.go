package main

type IDisplay interface {
	UpdateScanline(pixels []*Pixel, palette byte, ly int)
	HandleEvent() bool
	UpdateGraphics() error
}
