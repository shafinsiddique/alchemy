package main

import "image/color"

const (
	Z_FLAG                       = 7
	NEGATIVE_FLAG                = 6
	HALF_CARRY_FLAG              = 5
	CARRY_FLAG                   = 4
	SIXTEEN_BIT_INC_CYCLE        = 8
	LD_CYCLE                     = 4
	LD_SIXTEEN_BIT_CYCLE         = 16
	SCANLINE_CYCLES              = 456
	LY_INDEX                     = 0xFF44
	LCDC_INDEX                   = 0xFF40
	SCX_INDEX                    = 0xFF43
	SCY_INDEX                    = 0xFF42
	NUMBER_OF_TILES_IN_SCANLINE  = 20
	NUMBER_OF_PIXELS_IN_SCANLINE = 160
	NUMBER_OF_PIXELS_IN_TILE     = 8
	BGP_INDEX                    = 0xFF47
	TITLE                        = "Alchemist : A Game Boy Emulator"
	WINDOW_WIDTH                 = 160
	WINDOW_HEIGHT                = 144
)

var COLOR_MAP = map[byte]color.RGBA{
	3: {R: 0, G: 0, B: 0},
	2: {R: 169, G: 169, B: 169},
	1: {R: 211, G: 211, B: 211},
	0: {R: 255, G: 255, B: 255},
}
