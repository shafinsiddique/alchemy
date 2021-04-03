package main

import "fmt"

type CPU struct {
	Registers *Registers
	Memory []byte
	PC uint16
}

func NewCPU() *CPU {
	return &CPU{Registers: InitializeRegisters(), Memory: make([]byte,0x10000)}
}

func (cpu *CPU) LoadBootRom(bootrom []byte) {
	for i := 0; i < len(bootrom) ; i++ {
		hex := fmt.Sprintf("%x", bootrom[i])
		fmt.Println("0x" + hex)
		cpu.Memory[i] = bootrom[i]
	}
}

