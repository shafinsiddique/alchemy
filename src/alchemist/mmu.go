package main

type MMU struct {
	BootRomMemory []byte
	Memory        []byte
	BootMode      bool
	Joypad 		*byte
	Count int
	Debug *bool
}

func NewMMU(joypad *byte) *MMU {
	return &MMU{BootRomMemory: make([]byte, 256), Memory: make([]byte, 0x10000), Joypad: joypad}
}

func (mmu *MMU) Read(addr uint16) byte {
	if mmu.BootMode && addr < 256 {
		return mmu.BootRomMemory[addr]
	} else if addr == JOYPAD_INDEX {
		return mmu.getJoypadState()
	}

	return mmu.Memory[addr]
}

func (mmu *MMU) getJoypadState() byte {
	mmu.Count += 1
	val := mmu.Memory[JOYPAD_INDEX]
	val = SetBit(val, 1, 7)
	val = SetBit(val, 1, 6)

	var starting int
	if GetBit(val, 4) == 0 {
		starting = 0
	} else {
		starting = 4
	}

	for i := 0; i<4; i++ {
		status := GetBit(*mmu.Joypad, starting+i)
		//if status != 1 {
		//	*mmu.Debug = true
		//}
		val = SetBit(val, status, i)
	}

	return val
}

func (mmu *MMU) inRomRegion(addr uint16) bool {
	if addr <= 0x3fff || (addr >= 0x4000 && addr <= 0x7fff) {
		return true
	}
	return false
}

func (mmu *MMU) Write(addr uint16, val byte) {
	if mmu.BootMode && addr < 256 {
		mmu.BootRomMemory[addr] = val
	} else if mmu.BootMode {
		mmu.Memory[addr] = val
	} else if addr == JOYPAD_INDEX {
		mmu.writeToJoypad(val)
	} else if addr == DMA_INDEX {
		mmu.dmaTransfer(val)
	} else {
		if !mmu.inRomRegion(addr) {
			mmu.Memory[addr] = val
		}
	}
}

func (mmu *MMU) dmaTransfer(val byte) {
	src := uint16(val) << 8
	for i := 0; i < 160; i++ {
		index := uint16(i)
		obj := mmu.Read(src+index)
		mmu.Write(OAM_START+index, obj)
	}
}

func (mmu *MMU) writeToJoypad(val byte) {
	val = SetBit(val, GetBit(val, 4), 4) // we set only bit 4 and bit 5.
	val = SetBit(val, GetBit(val, 5), 5)
	mmu.Memory[JOYPAD_INDEX] = val
}

func (mmu *MMU) SetBootMode() {
	mmu.BootMode = true
}

func (mmu *MMU) SetRegularMode() {
	mmu.BootMode = false
}

func (mmu *MMU) LoadBootRom(bootrom []byte) {
	mmu.SetBootMode()
	for i := 0; i < len(bootrom); i++ {
		mmu.Write(uint16(i), bootrom[i])
	}
	mmu.SetRegularMode()
}

func (mmu *MMU) LoadBankRom0(rom []byte) {
	mmu.SetRegularMode()
	for i := 0; i < len(rom); i++ {
		mmu.Memory[uint16(i)] = rom[i]
	}
}
