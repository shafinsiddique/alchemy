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


	ppu.Registers.LY.Increment()
}

func (ppu *PPU) UpdateGraphics(cycles int) {
	if ppu.LCDEnabled() {
		ppu.Cycles -= cycles
		if ppu.Cycles <= 0 { // new scanline.
			ppu.StartScanline()
		}
	}
}

func (ppu *PPU) LCDEnabled() bool {
	return ppu.Registers.LCDC.GetBit(7) == 1
}



