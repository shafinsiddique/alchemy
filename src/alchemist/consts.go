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
	TITLE                        = "alchemist"
	WINDOW_WIDTH                 = 160
	WINDOW_HEIGHT                = 144
	MAX_CYCLES                   = 69905
	IE_INDEX                     = 0xFFFF
	IF_INDEX                     = 0xFF0F
	V_BLANK                      = 0
	LCD_STAT                     = 1
	TIMER                        = 2
	SERIAL                       = 3
	JOYPAD                       = 4
	JOYPAD_INDEX = 0xFF00
	RIGHT_JOYPAD = 0
	LEFT_JOYPAD = 1
	UP_JOYPAD = 2
	DOWN_JOYPAD = 3
	A_JOYPAD = 4
	B_JOYPAD = 5
	SELECT_JOYPAD = 6
	START_JOYPAD = 7
)

var COLOR_MAP = map[byte]color.RGBA{
	3: {R: 0, G: 0, B: 0},
	2: {R: 169, G: 169, B: 169},
	1: {R: 211, G: 211, B: 211},
	0: {R: 255, G: 255, B: 255},
}

var INTERRUPT_INSTRUCTIONS = map[byte]bool{0x0:false,0x1:false,0x76:false,0x2f:false,0x3f:false,
	0x27:false,0x37:false,0xf3:false,0xfb:false}
