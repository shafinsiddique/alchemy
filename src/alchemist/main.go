package main

import (
	"fmt"
	"os"
)

func RunBootSequence(cpu *CPU, mmu *MMU, ppu *PPU) {
	debug := false
	mmu.SetBootMode()
	for cpu.PC < 256 {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
		if cpu.PC == 0x006A {
			debug = true
		}

		if debug {
			fmt.Println(fmt.Sprintf("AF : %x", cpu.Registers.AF.Get()))
			fmt.Println(fmt.Sprintf("BC : %x", cpu.Registers.BC.Get()))
			fmt.Println(fmt.Sprintf("DE : %x", cpu.Registers.DE.Get()))
			fmt.Println(fmt.Sprintf("HL : %x", cpu.Registers.HL.Get()))
			fmt.Println(fmt.Sprintf("SP : %x", cpu.SP))
			fmt.Println(fmt.Sprintf("PC : %x", cpu.PC))
			fmt.Println(fmt.Sprintf("LY : %x", ppu.Registers.LY.Get()))
			fmt.Println(fmt.Sprintf("Cycles : %x", cycles))
			fmt.Println("END.")
		}
	}
}

func main() {
	mmu := NewMMU()
	cpu := &CPU{MMU: mmu, Registers: InitializeRegisters()}
	ppu := &PPU{Registers: InitializePPURegisters(mmu.Memory), Cycles: SCANLINE_CYCLES}
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/bootstrap.gb")
	f2, _ := os.Open(p + "/tetris.gb")
	bootrom := make([]byte, 256)
	rom := make([]byte, 0x10000)
	read, _ := f.Read(bootrom)
	romRead, _ := f2.Read(rom)
	fmt.Println("Bytes Read", read)
	fmt.Println("Rom Bytes Read", romRead)
	mmu.LoadBootRom(bootrom)
	mmu.LoadBankRom0(rom)
	RunBootSequence(cpu, mmu, ppu)

}
