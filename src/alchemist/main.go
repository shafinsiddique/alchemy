package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
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
	cyclesThisUpdate := 0
	for cyclesThisUpdate < MAX_CYCLES && cpu.PC < 256 {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
		cyclesThisUpdate += cycles
	}
}

func main() {
	initLogger()
	mmu := NewMMU()
	cpu := &CPU{MMU: mmu, Registers: InitializeRegisters()}
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/bootstrap.gb")
	f2, _ := os.Open(p + "/tetris.gb")
	bootrom := make([]byte, 256)
	rom := make([]byte, 0x10000)
	_, _ = f.Read(bootrom)
	_, _ = f2.Read(rom)
	//fmt.Println("Bytes Read", read)
	//fmt.Println("Rom Bytes Read", romRead)
	mmu.LoadBootRom(bootrom)
	mmu.LoadBankRom0(rom)
	mmu.SetBootMode()
	f.Close()
	f2.Close()
	dis, _ := NewSDLDisplay()
	ppu := &PPU{Registers: InitializePPURegisters(mmu.Memory), Cycles: SCANLINE_CYCLES, MMU: mmu, Display: dis}
	for dis.HandleEvent() {
		RunBootSequence(cpu, mmu, ppu)
		if cpu.PC >= 256 {
			//panic("end of prog.")
		} else {
			dis.UpdateGraphics()
			time.Sleep(1*time.Millisecond)
			//time.Sleep(2*time.Millisecond)
		}
	}
}
