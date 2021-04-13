package main

type MMU struct {
	BootRomMemory []byte
	Memory        []byte
	BootMode      bool
}

func NewMMU() *MMU {
	return &MMU{BootRomMemory: make([]byte, 256), Memory: make([]byte, 0x10000)}
}

func (mmu *MMU) Read(addr uint16) byte {
	if mmu.BootMode && addr < 256 {
		return mmu.BootRomMemory[addr]
	}

	return mmu.Memory[addr]
}

func (mmu *MMU) Write(addr uint16, val byte) {
	if mmu.BootMode && addr < 256 {
		mmu.BootRomMemory[addr] = val
	} else {
		mmu.Memory[addr] = val
	}
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
		mmu.Write(uint16(i), rom[i])
	}
}
