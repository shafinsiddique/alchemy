package main

type PPU struct {
	MMU       *MMU
	Registers *PPURegisters
	Cycles    int
}

func (ppu *PPU) StartScanline(){
	ppu.Cycles = SCANLINE_CYCLES
	// OAM Search.
	// Drawing
	// H-Blank
	// V-Blank.
}
func (ppu *PPU) UpdateGraphics(cycles int) {
	if ppu.LCDEnabled() {
		currentCycles := ppu.Cycles - cycles

		if currentCycles <= 0 { // new scanline.
			ppu.Registers.LY.Increment()
			ppu.StartScanline()
		}
	}
}

func (ppu *PPU) LCDEnabled() bool {
	return ppu.Registers.LCDC.GetBit(7) == 1
}



