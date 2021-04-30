package main

import (
	"math"
)

type PPU struct {
	MMU       *MMU
	Registers *PPURegisters
	Cycles    int
	Display   IDisplay
}

func NewPPU(mmu *MMU, display IDisplay) *PPU{
	return &PPU{Registers: InitializePPURegisters(mmu.Memory), Cycles: SCANLINE_CYCLES, MMU: mmu,
		Display: display}
}

func (ppu *PPU) IncrementScanline() {
	ppu.Registers.LY.Increment()
	if line := ppu.Registers.LY.Get(); line == 144 {
		ppu.Registers.IF.SetBit(1, V_BLANK)
	} else if line == 154 {
		ppu.Registers.LY.Set(0)
	}

}

func (ppu *PPU) FetchPixels() []*Pixel {
	firstTileAddr, lineNumber := ppu.getFirstTileIdAddr()
	row := make([]*Pixel, NUMBER_OF_PIXELS_IN_SCANLINE)
	offset := ppu.getFirstOffset() // we need an offset for the first one and then zero everytime afterwards.
	pixelCount := 0

	//found := false
	for i := 0; i < NUMBER_OF_TILES_IN_SCANLINE; i++ {
		tileId := ppu.MMU.Read(firstTileAddr + uint16(i))
		pixels := ppu.getHorizontalPixelsFromTile(tileId, lineNumber)
		for offset < NUMBER_OF_PIXELS_IN_TILE && pixelCount < NUMBER_OF_PIXELS_IN_SCANLINE {
			row[pixelCount] = pixels[offset]
			pixelCount += 1
			offset += 1
		}
		offset = 0
	}

	return row
}

func (ppu *PPU) getFirstOffset() int {
	sx := int(ppu.Registers.SCX.Get())
	return sx % 8
}

func (ppu *PPU) getHorizontalPixelsFromTile(tileId byte, lineNumber uint16) []*Pixel {
	addr := ppu.getTileAddr(tileId)
	lineStartingIndex := addr + (lineNumber * 2)
	low := ppu.MMU.Read(lineStartingIndex)
	high := ppu.MMU.Read(lineStartingIndex + 1)
	return GetPixels(high, low)
}

func (ppu *PPU) getTileAddr(tileId byte) uint16 {
	var addr uint16
	tileNo := uint16(tileId)
	if ppu.Registers.LCDC.GetBit(4) == 1 {
		addr = 0x8000 + (tileNo*16)
	} else {
		complement, isNegative := GetTwosComplement(tileId)
		offset := uint16(complement) * 16
		if isNegative {
			addr -= offset
		} else {
			addr += offset
		}
	}

	return addr
}

func (ppu *PPU) getFirstTileIdAddr() (uint16, uint16) {
	totalYOffsetInPixels := ppu.Registers.SCY.Get() + ppu.Registers.LY.Get()
	tileBlockStartingAddr :=  ppu.getBackgroundMapAddr() + (uint16(math.Floor(float64(totalYOffsetInPixels)/8)*32))
	scanlineOffset := uint16(totalYOffsetInPixels % 8)
	withXOffset := tileBlockStartingAddr + uint16(math.Floor(float64(ppu.Registers.SCX.Get())/8))
	return withXOffset, scanlineOffset
}

func (ppu *PPU) getBackgroundMapAddr() uint16 {
	return 0x9800 // need to determine based on register.
}

func (ppu *PPU) runScanline() {
	ppu.Cycles = SCANLINE_CYCLES
	// OAM Search.
	// Drawing
	if ppu.Registers.LY.Get() <= 143 {
		pixels := ppu.FetchPixels()
		ppu.Display.UpdateScanline(pixels, ppu.Registers.BGP.Get(), int(ppu.Registers.LY.Get()))
	} // else it is in H_BLANK.

	// H-Blank
	// V-Blank.

	ppu.IncrementScanline()
}
func (ppu *PPU) handleDisabledLCD() {
	ppu.Cycles = SCANLINE_CYCLES
	ppu.Registers.LY.Set(0)
	ppu.setModeInLCDStat(0)
}

func (ppu *PPU) UpdateGraphics(cycles int) {
	if !ppu.LCDEnabled() {
		ppu.handleDisabledLCD()
		return
	}
	ppu.setLCDStatus()
	ppu.Cycles -= cycles
	if ppu.Cycles <= 0 { // new scanline. 456 Clock Cycles has past, so we have a new scanline.
		ppu.runScanline()
	}
}

func (ppu *PPU) setModeInLCDStat(mode byte) {
	b1, b0 := GetBit(mode, 1), GetBit(mode, 0)
	ppu.Registers.LCD_STAT.SetBit(b1, 1)
	ppu.Registers.LCD_STAT.SetBit(b0, 0)
}

func  (ppu *PPU) getCurrentMode() byte {
	var mode byte
	lcdStat := ppu.Registers.LCD_STAT
	mode = SetBit(mode, lcdStat.GetBit(1),1)
	mode = SetBit(mode, lcdStat.GetBit(0), 0)
	return mode
}

func (ppu *PPU) checkForInterrupts(oldMode byte, newMode byte) {
	mode := int(newMode)
	if (newMode == 0 || newMode == 1 || newMode == 2) && (oldMode != newMode) &&
		ppu.Registers.LCD_STAT.GetBit(mode+3) == 1 {
		ppu.requestPPUInterrupt()
	}
}

func (ppu *PPU) requestPPUInterrupt(){
	ppu.Registers.IF.SetBit(1, LCD_STAT)
}

func (ppu *PPU) checkForCoincidence() {
	var bit byte = 0
	if ppu.Registers.LY.Get() == ppu.MMU.Read(0xff45) {
		bit = 1
		if ppu.Registers.LCD_STAT.GetBit(6) == 1 {
			ppu.requestPPUInterrupt()
		}
	}

	ppu.Registers.LCD_STAT.SetBit(bit, 2)
}

func (ppu *PPU) setLCDStatus() {
	line := ppu.Registers.LY.Get()
	currentMode := ppu.getCurrentMode()
	var mode byte
	if line >= 144 {
		mode = 1
	} else {
		if ppu.Cycles >= SCANLINE_CYCLES-MODE_2_CYCLES {
			mode = 2
		} else if ppu.Cycles >= SCANLINE_CYCLES-MODE_2_CYCLES-MODE_3_CYCLES {
			mode = 3
		} else {
			mode = 0
		}
	}
	ppu.setModeInLCDStat(mode)
	ppu.checkForInterrupts(currentMode, mode)
	ppu.checkForCoincidence()

}

func (ppu *PPU) LCDEnabled() bool {
	return ppu.Registers.LCDC.GetBit(7) == 1
}
