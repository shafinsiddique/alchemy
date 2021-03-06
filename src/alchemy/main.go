package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func initializeComponents() (*CPU, *MMU, *PPU, IDisplay) {
	joypad := byte(0b11111111)
	mmu := NewMMU(&joypad)
	cpu := NewCPU(mmu)
	display, _ := NewSDLDisplay(&joypad, cpu.Registers.IF)
	display.IE = cpu.Registers.IE
	display.CPU = cpu
	ppu := NewPPU(mmu, display)
	return cpu, mmu, ppu, display
}

func load(bootrom string, rom string, mmu *MMU) {
	bootromFile, err := os.Open(bootrom)
	if err != nil {
		log.Fatal("Unable to load boot rom.")
	}
	romFile, err := os.Open(rom)
	if err != nil {
		log.Fatal("Unable to load rom.")
	}
	bRomArr := make([]byte, 256)
	romArr := make([]byte, 0x10000)
	_, _ = bootromFile.Read(bRomArr)
	_, _ = romFile.Read(romArr)
	mmu.LoadBootRom(bRomArr[:])
	mmu.LoadBankRom0(romArr[:])
_:
	bootromFile.Close()
_:
	romFile.Close()
}

func run(cpu *CPU, mmu *MMU, ppu *PPU) {
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MAX_CYCLES {
		opcode := cpu.FetchAndIncrement()
		cycles := cpu.Execute(opcode)
		cpu.UpdateTimers(cycles)
		ppu.UpdateGraphics(cycles)
		cpu.HandleInterrupts(opcode)

		cyclesThisUpdate += cycles
		if mmu.BootMode && cpu.PC >= 256 {
			mmu.SetRegularMode()
		}
	}
}

func main() {
	cpu, mmu, ppu, dis := initializeComponents()
	p, _ := os.Getwd()
	load(p+"/bootstrap.gb", p+"/tetris.gb", mmu)
	mmu.SetBootMode()
	for dis.HandleEvent() {
		run(cpu, mmu, ppu)
		_ = dis.UpdateGraphics()
		time.Sleep(10 * time.Millisecond)
	}

}
