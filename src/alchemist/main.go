package main

import (
	"fmt"
	"os"
)

func main() {
	cpu := NewCPU()
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/src/alchemist/bootstrap.gb")
	bootrom := make([]byte, 256)
	read, _ := f.Read(bootrom)
	fmt.Println("Bytes Read", read)
	cpu.LoadBootRom(bootrom)
}
