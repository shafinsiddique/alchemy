package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func run(cpu *CPU, mmu *MMU, ppu *PPU) {
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MAX_CYCLES {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
		cyclesThisUpdate += cycles
		if mmu.BootMode && cpu.PC >= 256 {
			log.Fatal("end of prog.")
			mmu.SetRegularMode()
		}
	}
}

func main() {
	mmu := NewMMU()
	cpu := &CPU{MMU: mmu, Registers: InitializeRegisters()}
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/bootstrap.gb")
	f2, _ := os.Open(p + "/tetris.gb")
	bootrom := make([]byte, 256)
	rom := make([]byte, 0x10000)
	_, _ = f.Read(bootrom)
	_, _ = f2.Read(rom)
	mmu.LoadBootRom(bootrom)
	mmu.LoadBankRom0(rom)
	f.Close()
	f2.Close()
	dis, _ := NewSDLDisplay()
	ppu := &PPU{Registers: InitializePPURegisters(mmu.Memory), Cycles: SCANLINE_CYCLES, MMU: mmu, Display: dis}
	mmu.SetBootMode()
	for dis.HandleEvent() {
		run(cpu, mmu, ppu)
		dis.UpdateGraphics()
		time.Sleep(10*time.Millisecond)

	}

}
