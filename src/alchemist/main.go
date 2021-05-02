package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func initializeComponents() (*CPU, *MMU, *PPU, IDisplay) {
	debug := false
	joypad := byte(0b11111111)
	mmu := NewMMU(&joypad)
	cpu := NewCPU(mmu)
	display, _ := NewSDLDisplay(&joypad, cpu.Registers.IF)
	display.IE = cpu.Registers.IE
	display.CPU = cpu
	ppu := NewPPU(mmu, display)
	cpu.Debug = &debug
	mmu.Debug = &debug
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
	debug2 := false
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MAX_CYCLES {
		opcode := cpu.FetchAndIncrement()
		cycles := cpu.Execute(opcode)
		ppu.UpdateGraphics(cycles)
		cpu.HandleInterrupts(opcode)
		if *cpu.Debug && cpu.PC == 0x29A6 {
			debug = true
		}

		if !mmu.BootMode && cpu.PC == 0x40 {
			debug2 = true
		}
		//if GetBit(*mmu.Joypad, 6) == 0 {
		//	debug = true
		//}
		//
		//if debug {
		//	if cpu.PC > 0x2f0 {
		//		debug2 = true
		//	}
		//}
		//if cpu.PC == 0x40 && cpu.Registers.DE.Get() == 0x0479 {
		//	debug = true
		//}
		//
		//if debug {
		//	if cpu.PC == 0x02ED {
		//		debug2 = true
		//	}
		//}
		//
		//if debug2 {
		//	if cpu.PC > 0x02F0 {
		//		//fmt.Println(fmt.Sprintf("%x A:", cpu.Registers.A.Get()))
		//		//fmt.Println(fmt.Sprintf("%x LY:", ppu.Registers.LY.Get()))
		//
		//	}
		//}
		//if !mmu.BootMode && mmu.Read(0xff02) == 0x81{
		//	fmt.Println(fmt.Sprintf("%c", mmu.Read(0xff01)))
		//	mmu.Memory[0xff02] = 0
		//}
		//if !mmu.BootMode && cpu.PC == 0xC2b6 {
		//	debug = true
		//}


		if  debug {
			fmt.Println(fmt.Sprintf("%x", mmu.Read(0xff80)))
			fmt.Println(fmt.Sprintf("AF: %x", cpu.Registers.AF.Get()))
			fmt.Println(fmt.Sprintf("BC: %x", cpu.Registers.BC.Get()))
			fmt.Println(fmt.Sprintf("DE: %x", cpu.Registers.DE.Get()))
			fmt.Println(fmt.Sprintf("HL: %x", cpu.Registers.HL.Get()))
			fmt.Println(fmt.Sprintf("SP: %x", cpu.SP))
			fmt.Println(fmt.Sprintf("PC: %x", cpu.PC))
			fmt.Println("END.")
		}

		if debug2{
			//fmt.Println(fmt.Sprintf("AF: %x", cpu.Registers.AF.Get()))
			//fmt.Println(fmt.Sprintf("BC: %x", cpu.Registers.BC.Get()))
			//fmt.Println(fmt.Sprintf("DE: %x", cpu.Registers.DE.Get()))
			//fmt.Println(fmt.Sprintf("HL: %x", cpu.Registers.HL.Get()))
			//fmt.Println(fmt.Sprintf("SP: %x", cpu.SP))
			//fmt.Println(fmt.Sprintf("PC: %x", cpu.PC))
			//fmt.Println("END.")
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
		_ = dis.UpdateGraphics()
		time.Sleep(10*time.Millisecond)
	}
}
