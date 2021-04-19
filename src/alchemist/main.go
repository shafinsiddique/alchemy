package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
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
	mmu.SetBootMode()
	for cpu.PC < 256 {
		cycles := cpu.FetchDecodeExecute()
		ppu.UpdateGraphics(cycles)
	}
	mmu.SetRegularMode()
}

func main() {
	//initLogger()
	//mmu := NewMMU()
	//cpu := &CPU{MMU: mmu, Registers: InitializeRegisters()}
	//ppu := &PPU{Registers: InitializePPURegisters(mmu.Memory), Cycles: SCANLINE_CYCLES, MMU: mmu}
	//p, _ := os.Getwd()
	//f, _ := os.Open(p + "/bootstrap.gb")
	//f2, _ := os.Open(p + "/tetris.gb")
	//bootrom := make([]byte, 256)
	//rom := make([]byte, 0x10000)
	//read, _ := f.Read(bootrom)
	//romRead, _ := f2.Read(rom)
	//fmt.Println("Bytes Read", read)
	//fmt.Println("Rom Bytes Read", romRead)
	//mmu.LoadBootRom(bootrom)
	//mmu.LoadBankRom0(rom)
	//RunBootSequence(cpu, mmu, ppu)
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	//defer sdl.Quit()

	window, err := sdl.CreateWindow("Alchemist : A Gameboy Emulator",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		160, 144, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	//defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	c := color.RGBA{R:255, G:255, B:255}
	for y := 0; y<144; y++ {
		for x := 0; x<160; x++ {
			surface.Set(x, y, c)
		}
	}
	window.UpdateSurface()

	//running := true
	//for running {
	//	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	//		switch event.(type) {
	//		case *sdl.QuitEvent:
	//			running = false
	//			break
	//		}
	//	}
	//}
	fmt.Println(MergeBytes(1,1))

}
