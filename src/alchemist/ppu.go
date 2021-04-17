package main

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

func (ppu *PPU) FetchPixels(){
	// Algorithm : SCY -> How many pixels from the top of the background is the viewport.
	// SCX -> how many pixels from the left is the viewport.
	// First , we need to get the TILE ID. We essentially NEED 20 tiles for this scanline.
	// we have an sx, sy. we have a scanline.
	// in background map, tile ids are stored in a 32 x 32 grid. 0-32 is the first scanline.
	// 33-.. is the second scanline and so on.
	// We use sy to compute which region of 32 block our tile will be in.
	// region := math.Floor(sy/32).
	// Now we have the region, we need the starting index of that block.
	// blockStartingIndex := region * 32 -> 0*32 => 0, 1*32 -> 32.
	// Once we have the starting index of the block, we essentially know what row we are on.
	// now we need to compute from that row, which column are we on.
	// sx -> How many pixels from the left are we at.
	// each byte in the background map is 8 PIXELS. we need to know which byte from the region we are at.
	// tileByte := blockStartingIndex + (sx/8) (sx/8 because 1 pixel from the left would be in the 0th tile. therefore
	// blockstartingindex + 0. 8 pixels in we'd be at the 1th tile, therefore addr := blockingstartingindex + 1.

	// so now we know EXACTLY which from which tile we're goign to be startign at. What we need to do is.
	// compute from WHICH pixel of that tile, rememeber there are 8 in each tile, do we start our SCANLINE from.
	// We have our tyleByte and can get the tileId:

	// tileId := mmu.Read[tyleByte]
	// we can go into tile ram and fetch that tile.
	// tile := mmu.Read[getTireAddrr(tileId)]
	// WHich horizontal line are we at?
	// pixels := mmu.Read[tile + scanline % 8] => merge with second bytes.
	// get pixels back.
	// do we need to truncate?


}

func (ppu *PPU) StartScanline() {
	ppu.Cycles = SCANLINE_CYCLES
	// OAM Search.
	// Drawing
	ppu.FetchPixels()
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
