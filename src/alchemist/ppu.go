package main

import (
	"fmt"
	"math"
)

type PPU struct {
	MMU       *MMU
	Registers *PPURegisters
	Cycles    int
}

func (ppu *PPU) IncrementScanline() {
	ppu.Registers.LY.Increment()
	if ppu.Registers.LY.Get() == 154 {
		ppu.Registers.LY.Set(0)
	}
}

func (ppu *PPU) FetchPixels() []*Pixel {
	if ppu.Registers.SCY.Get() < 98 {
		fmt.Println(ppu.Registers.SCY.Get())
		panic("end of prog")
	}
	firstTileAddr := ppu.getFirstTileIdAddr()
	row := make([]*Pixel, NUMBER_OF_PIXELS_IN_SCANLINE)
	offset := ppu.getFirstOffset() // we need an offset for the first one and then zero everytime afterwards.
	pixelCount := 0
	for i := 0; i < NUMBER_OF_TILES_IN_SCANLINE; i++ {
		tileId := ppu.MMU.Read(firstTileAddr+uint16(i))

		pixels := ppu.getHorizontalPixelsFromTile(tileId)
		for offset < 7 && pixelCount < NUMBER_OF_PIXELS_IN_SCANLINE {
			row[pixelCount] = pixels[offset]
			pixelCount += 1
		}
		offset = 0
	}

	return row

}

func (ppu *PPU) getFirstOffset() int {
	sx := int(ppu.Registers.SCX.Get())
	return sx - int(math.Floor(float64(sx/8))*8)
}

func (ppu *PPU) getHorizontalPixelsFromTile(tileId byte) []*Pixel {
	addr := ppu.getTileAddr(tileId)
	lineNumber := uint16(ppu.Registers.LY.Get() % 8)
	lineStartingIndex := addr + (lineNumber * 2)
	low := ppu.MMU.Read(lineStartingIndex)
	high := ppu.MMU.Read(lineStartingIndex + 1)
	return GetPixels(high, low)
}

func (ppu *PPU) getTileAddr(tileId byte) uint16 {
	var addr uint16
	if ppu.Registers.LCDC.GetBit(4) == 1 {
		addr = 0x8000 + uint16(tileId*16)
	} else {
		complement, isNegative := GetTwosComplement(tileId)
		offset := uint16(complement)*16
		if isNegative {
			addr -= offset
		} else {
			addr += offset
		}
	}

	return addr
}

func (ppu *PPU) getFirstTileIdAddr() uint16 {
	//tileBlockStartingIndex := (math.Floor(float64(ppu.Registers.SCY.Get()/32)) * 32) +
	//	float64(ppu.getBackgroundMapAddr()) // the section in background map where the tile is.
	//
	//tileBlockOffset := math.Floor((float64(ppu.Registers.SCX.Get())) / 8) // how many tiles we have to skip
	// horizontally.
	//
	tileBlockStartingAddr := float64(ppu.getBackgroundMapAddr()) + (math.Floor(float64(ppu.Registers.SCY.Get()/8))*32)
	tileBlockOffsetY := tileBlockStartingAddr + (math.Floor(float64(ppu.Registers.LY.Get()/8))*32)
	tileBlockOffsetX := math.Floor(float64(ppu.Registers.SCX.Get() / 8))
	return uint16(tileBlockOffsetY + tileBlockOffsetX)
}

func (ppu *PPU) getBackgroundMapAddr() uint16 {
	return 0x9800 // need to determine based on register.
}

func (ppu *PPU) StartScanline() {
	ppu.Cycles = SCANLINE_CYCLES
	// OAM Search.
	// Drawing
	pixels := ppu.FetchPixels()
	//if pixels[0].high != 0 || pixels[0].low != 0 {
	//	fmt.Println("hello")
	//}
	for i := 0; i < len(pixels); i++  {
		if pixels[i].high != 0 || pixels[i].low != 0 {
			fmt.Println("hello")
		}
	}
	// H-Blank
	// V-Blank.

	ppu.IncrementScanline()
}

func (ppu *PPU) UpdateGraphics(cycles int) {
	if !ppu.LCDEnabled() {
		return
	}

	ppu.Cycles -= cycles
	if ppu.Cycles <= 0 { // new scanline.
		ppu.StartScanline()
	}
}

func (ppu *PPU) LCDEnabled() bool {
	return ppu.Registers.LCDC.GetBit(7) == 1
}
