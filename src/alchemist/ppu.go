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

func (ppu *PPU) incrementScanline() {
	ppu.Registers.LY.Increment()
	if line := ppu.Registers.LY.Get(); line == 144 {
		ppu.Registers.IF.SetBit(1, V_BLANK)
	} else if line == 154 {
		ppu.Registers.LY.Set(0)
	}
}

func (ppu *PPU) spriteIsVisible(x byte, y byte, spriteHeight byte) bool {
	xPos := int(x)
	yPos := int(y)
	ly := int(ppu.Registers.LY.Get())
	height := int(spriteHeight)
	if xPos - 8 > -8 && yPos - 16 >= ly && ly <= yPos + height{
		return true
	}

	return false
}

func getSpriteHeight() byte {
	return 8 // need to add logic for 16s.
}

func (ppu *PPU) addTileToPixels(pixels []*Pixel, tilePixels []*Pixel, x byte) {
	xPos := int(x)-8
	var startingInPixels byte
	var startingInTile byte
	if xPos < 0 {
		startingInTile = byte(math.Abs(float64(xPos)))
		startingInPixels = 0
	} else {
		startingInTile = 0
		startingInPixels = x-8
	}

	for startingInTile < NUMBER_OF_PIXELS_IN_TILE && startingInPixels < NUMBER_OF_PIXELS_IN_SCANLINE {
		pixels[startingInPixels] = tilePixels[startingInTile]
		startingInTile += 1
		startingInPixels += 1
	}
}

func (ppu *PPU) fetchSprites() []*Sprite {
	sprites := make([]*Sprite, 10)
	xValues := map[byte]bool{}
	spriteHeight := getSpriteHeight()
	count := 0
	for i := 0xFE00; i < 0xFE9f; i+=4 {
		index := uint16(i) // need to convert for signature
		sprite := NewSprite(ppu.MMU.Read(index), ppu.MMU.Read(index+1), ppu.MMU.Read(index+2),
			ppu.MMU.Read(index+3))

		if _, ok := xValues[sprite.X];
		ppu.spriteIsVisible(sprite.X, sprite.Y, spriteHeight) && count < 10 && !ok{
			//tilePixels := ppu.getHorizontalPixelsFromTile(sprite.TileID, 9) // fix.
			//ppu.addTileToPixels(pixels, tilePixels, sprite.X)
			sprites[count] = sprite
			xValues[sprite.X] = true // go needs hash sets -_-.
			count += 1
		} else if count == 10 {
			break
		}
	}

	return sprites
}

func (ppu *PPU) fetchPixels() []*Pixel {
	firstTileAddr, lineNumber := ppu.getFirstTileIdAddr()
	row := make([]*Pixel, NUMBER_OF_PIXELS_IN_SCANLINE)
	offset := ppu.getFirstOffset() // we need an offset for the first one and then zero everytime afterwards.
	pixelCount := 0

	//found := false
	for i := 0; i < NUMBER_OF_TILES_IN_SCANLINE; i++ {
		tileId := ppu.MMU.Read(firstTileAddr + uint16(i))
		pixels := ppu.getHorizontalPixelsFromTile(tileId, lineNumber, false)
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

func (ppu *PPU) getHorizontalPixelsFromTile(tileId byte, lineNumber uint16, flipped bool) []*Pixel {
	var high byte
	var low byte
	addr := ppu.getTileAddr(tileId)
	if flipped {
		addr += 15
		lineStartingIndex := addr - (lineNumber * 2)
		high = ppu.MMU.Read(lineStartingIndex)
		low = ppu.MMU.Read(lineStartingIndex-1)

	} else {
		addr = ppu.getTileAddr(tileId)

		lineStartingIndex := addr + (lineNumber * 2)
		low = ppu.MMU.Read(lineStartingIndex)
		high = ppu.MMU.Read(lineStartingIndex + 1)
	}

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
		pixels := ppu.fetchPixels()
		ppu.Display.UpdateScanline(pixels, ppu.Registers.BGP.Get(), int(ppu.Registers.LY.Get()))
	} // else it is in H_BLANK.

	// H-Blank
	// V-Blank.

	ppu.incrementScanline()
}

func (ppu *PPU) handleDisabledLCD() {
	ppu.Cycles = SCANLINE_CYCLES
	ppu.Registers.LY.Set(0)
	ppu.setModeInLCDStat(0)
}

func (ppu *PPU) getSpriteTileAndLine(sprite *Sprite, height byte) (tileId byte, lineNumber uint16){
	yPos := int(sprite.Y) - 16
	if yPos < 0 {
		lineNumber = uint16(0 + ppu.Registers.LY.Get())
	} else {
		lineNumber = uint16(ppu.Registers.LY.Get() - (sprite.Y - 16))
	}

	if height == 8 {
		tileId = sprite.TileID

	} else {
		tileA, tileB := sprite.TileID & 0XFE, (sprite.TileID & 0xFe) + 1 // tileA always top, tileB bottom.
		if sprite.YFlip {tileB, tileA = tileA, tileB}

		tileId = tileA

		if lineNumber > 7 {
			tileId = tileB
			lineNumber = lineNumber % 8
		}
	}

	return tileId, lineNumber
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
