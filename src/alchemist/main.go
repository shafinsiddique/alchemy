package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func initializeComponents() (*CPU, *MMU, *PPU, IDisplay) {
	mmu := NewMMU()
	cpu := NewCPU(mmu)
	display, _ := NewSDLDisplay()
	ppu := NewPPU(mmu, display)
	return cpu, mmu, ppu, display
}

func load(bootrom string, rom string, mmu *MMU) {
	bootromFile, err := os.Open(bootrom)
	if err != nil {log.Fatal("Unable to load boot rom.")}
	romFile, err := os.Open(rom)
	if err != nil {log.Fatal("Unable to load rom.")}
	bRomArr := make([]byte, 256)
	romArr := make([]byte, 0x10000)
	_ ,_ = bootromFile.Read(bRomArr)
	_, _ = romFile.Read(romArr)
	mmu.LoadBootRom(bRomArr[:])
	mmu.LoadBankRom0(romArr[:])
	_ : bootromFile.Close()
	_ : romFile.Close()
}

func run(cpu *CPU, mmu *MMU, ppu *PPU) {
	debug := false
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MAX_CYCLES {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
		if cpu.PC == 0x2b8 {
			debug = true
		}
		if !mmu.BootMode && debug {
			fmt.Println(fmt.Sprintf("AF: %x", cpu.Registers.AF.Get()))
			fmt.Println(fmt.Sprintf("BC: %x", cpu.Registers.BC.Get()))
			fmt.Println(fmt.Sprintf("DE: %x", cpu.Registers.DE.Get()))
			fmt.Println(fmt.Sprintf("HL: %x", cpu.Registers.HL.Get()))
			fmt.Println(fmt.Sprintf("SP: %x", cpu.SP))
			fmt.Println(fmt.Sprintf("PC: %x", cpu.PC))
			fmt.Println("END.")
		}

		cyclesThisUpdate += cycles
		if mmu.BootMode && cpu.PC >= 256 {
			mmu.SetRegularMode()
		}
	}
}

func main() {
	cpu, mmu, ppu, dis := initializeComponents()
	p, _ := os.Getwd()
	load(p + "/bootstrap.gb", p + "/tetris.gb", mmu)
	mmu.SetBootMode()
	for dis.HandleEvent() {
		run(cpu, mmu, ppu)
		//_ = dis.UpdateGraphics()
		//time.Sleep(10*time.Millisecond)
	}

}
