package main

import "fmt"

type CPU struct {
	Registers *Registers
	Memory []byte
	PC uint16
	SP uint16
}

func NewCPU() *CPU {
	return &CPU{Registers: InitializeRegisters(), Memory: make([]byte,0x10000)}
}

func (cpu *CPU) LoadBootRom(bootrom []byte) {
	for i := 0; i < len(bootrom) ; i++ {
		cpu.Memory[i] = bootrom[i]
	}
}


func (cpu *CPU) FetchDecodeExecute() {
	pc := &cpu.PC
	switch opcode := cpu.Memory[*pc]; opcode {
	case 0x31:
		cpu.LD_SP_D16()
	case 0xaf:
		cpu.XOR_A()
	case 0x21:
		cpu.LD_HL_D16()
	case 0x32:
		cpu.LD_HL_A()
	case 0xcb:
		cpu.Oxcb()
	case 0x20:
		cpu.JR_NZ_S8() // s8 stands for signed 8 bit.
	default:
		hex := fmt.Sprintf("%x", opcode)
		fmt.Println("0x" + hex)
	}
	*pc += 1 // always increment one, even if other instructions increment, we need to increment from that position to
	// go to the next one. basically, this allows us to not worry abouyt incrementing one at the end of every single
	// function.

}

func (cpu *CPU) RunBootSequence(){
	for cpu.PC < 256 {
		cpu.FetchDecodeExecute()
	}
}

