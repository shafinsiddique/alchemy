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
	LCD_STATUS                   = 0xFF41
	SCX_INDEX                    = 0xFF43
	SCY_INDEX                    = 0xFF42
	NUMBER_OF_TILES_IN_SCANLINE  = 20
	NUMBER_OF_PIXELS_IN_SCANLINE = 160
	NUMBER_OF_PIXELS_IN_TILE     = 8
	BGP_INDEX                    = 0xFF47
	TITLE                        = "alchemy"
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
	JOYPAD_INDEX                 = 0xFF00
	RIGHT_JOYPAD                 = 0
	LEFT_JOYPAD                  = 1
	UP_JOYPAD                    = 2
	DOWN_JOYPAD                  = 3
	A_JOYPAD                     = 4
	B_JOYPAD                     = 5
	SELECT_JOYPAD                = 6
	START_JOYPAD                 = 7
	MODE_2_CYCLES                = 80
	MODE_3_CYCLES                = 172
	OBP0_INDEX                   = 0xFF48
	OBP1_INDEX                   = 0xFF49
	OBSCURE_COLOR                = 210
	TIMA_INDEX                   = 0xFF05
	TMA_INDEX                    = 0xff06
	TAC_INDEX                    = 0xff07
	DIV_INDEX                    = 0xff04
	LYC_INDEX                    = 0xff45
	OAM_START                    = 0xFE00
	OAM_END                      = 0xFE95
	DMA_INDEX                    = 0xFF46
)

var COLOR_MAP = map[byte]color.RGBA{
	3: {R: 0, G: 0, B: 0},       // black.
	2: {R: 136, G: 192, B: 112}, // Light Gray
	1: {R: 52, G: 104, B: 86},   // Dark Grey
	0: {R: 224, G: 248, B: 208}, // White
}

var INTERRUPT_INSTRUCTIONS = map[byte]bool{0x0: false, 0x1: false, 0x76: false, 0x2f: false, 0x3f: false,
	0x27: false, 0x37: false, 0xf3: false, 0xfb: false}
