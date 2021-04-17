package main

// pixels in the gameboy are stored in the format 2BPP -> 2 Bits Per Pixel.

type Pixel struct {
	low byte
	high byte
}

func (pixel *Pixel) GetHigh() byte {
	return pixel.high
}

func (pixel *Pixel) GetLow() byte {
	return pixel.low
}


