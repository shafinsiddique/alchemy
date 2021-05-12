package main

// pixels in the gameboy are stored in the format 2BPP -> 2 Bits Per Pixel.

type Pixel struct {
	low  byte
	high byte
}

func GetPixels(high byte, low byte) []*Pixel {
	pixels := make([]*Pixel, 8)
	for i := 0; i < 8; i++ {
		highBit := GetBit(high, 7-i)
		lowBit := GetBit(low, 7-i)
		pixels[i] = &Pixel{low: lowBit, high: highBit}
	}
	return pixels
}

func (pixel *Pixel) GetHigh() byte {
	return pixel.high
}

func (pixel *Pixel) GetLow() byte {
	return pixel.low
}

func (pixel *Pixel) Get() byte {
	high := pixel.high
	low := pixel.low
	switch high {
	case 1:
		switch low {
		case 1:
			return 3
		default:
			return 2
		}
	default:
		return low
	}
}
