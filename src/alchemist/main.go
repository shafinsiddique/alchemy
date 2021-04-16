package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.New()
	path, _ := os.Getwd()
	logFile, err := os.OpenFile(path+"/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err == nil {
		logger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	} else {
		logger.Error("error trying to initialize log file.")
	}
}

func RunBootSequence(cpu *CPU, mmu *MMU, ppu *PPU) {
	debug := false
	mmu.SetBootMode()
	for cpu.PC < 256 {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
		if cpu.PC == 0x00e0 {
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
			fmt.Println("END.")
		}
	}
}

func main() {
	initLogger()
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
