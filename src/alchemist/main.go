package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var count = 0
var logger *log.Logger
func initLogger() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger = log.New()
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetFormatter(customFormatter)
	path, _ := os.Getwd()
	fmt.Println(path)
	logFile, err := os.OpenFile(path+"/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err == nil {
		logger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	} else {
		logger.Error("error trying to initialize log file.")
	}
}

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
	ppu.Debug = &debug
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
		cpu.UpdateTimers(cycles)
		ppu.UpdateGraphics(cycles)
		cpu.HandleInterrupts(opcode)


		if *cpu.Debug && cpu.PC == 0x1bce && mmu.Read(0xffe1) == 0x0 {
			debug = true
		}
		
		if debug {
			//fmt.Println(fmt.Sprintf("Game Status: %x", mmu.Read(0xffe1)))
			//fmt.Println(fmt.Sprintf("Button Hit %x", mmu.Read(0xff81)))
			//fmt.Println(fmt.Sprintf("Button Down %x", mmu.Read(0xff80)))
			//fmt.Println(fmt.Sprintf("LY: %x", ppu.Registers.LY.Get()))
			//fmt.Println(fmt.Sprintf("AF: %x", cpu.Registers.AF.Get()))
			//fmt.Println(fmt.Sprintf("BC: %x", cpu.Registers.BC.Get()))
			//fmt.Println(fmt.Sprintf("DE: %x", cpu.Registers.DE.Get()))
			//fmt.Println(fmt.Sprintf("HL: %x", cpu.Registers.HL.Get()))
			//fmt.Println(fmt.Sprintf("SP: %x", cpu.SP))
			//fmt.Println(fmt.Sprintf("PC: %x", cpu.PC))
			//fmt.Println("END.")
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
	initLogger()
	//var lastOpcode byte
	cpu, mmu, ppu, dis := initializeComponents()
	p, _ := os.Getwd()
	load(p + "/bootstrap.gb", p + "/Dr. Mario (World).gb", mmu)
	mmu.SetBootMode()
	for dis.HandleEvent() {
		run(cpu, mmu, ppu)
		_ = dis.UpdateGraphics()
		time.Sleep(10*time.Millisecond)
	}

}
