package main

import (
	"fmt"
	"os"
)

func main() {
	cpu := NewCPU()
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/bootstrap.gb")
	f2, _ := os.Open(p + "/tetris.gb")
	bootrom := make([]byte, 256)
	rom := make([]byte, 0x1000)
	read, _ := f.Read(bootrom)
	romRead, _ := f2.Read(rom)
	fmt.Println("Bytes Read", read)
	fmt.Println("Rom Bytes Read", romRead)
	cpu.LoadBootRom(bootrom)
	cpu.LoadRomBank0(rom)
	fmt.Println(fmt.Sprintf("%x",cpu.Memory[260]))
	//cpu.RunBootSequence()

}
