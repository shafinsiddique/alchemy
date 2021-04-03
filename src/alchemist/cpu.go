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
		hex := fmt.Sprintf("%x", bootrom[i])
		fmt.Println("0x" + hex)
		cpu.Memory[i] = bootrom[i]
	}
}

func (cpu *CPU) Ox31(){
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.SP = MergeBytes(high, low)
	*pc += 1
}
func (cpu *CPU) Oxaf() {
	pc := &cpu.PC
	*pc += 1
	cpu.Registers.A.Set(cpu.Registers.A.Value ^ cpu.Registers.A.Value)
}

func (cpu *CPU) Ox32() {
	pc := &cpu.PC
	a := cpu.Registers.A.Get()
	cpu.Memory[cpu.Registers.HL.Get()] = a
	cpu.Registers.HL.Decrement()
	*pc += 1
}


func (cpu *CPU) Ox21() {
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	*pc += 1
	cpu.Registers.H.Set(high)
	cpu.Registers.L.Set(low)
}

func (cpu *CPU) FetchDecodeExecute() {
	pc := cpu.PC
	switch opcode := cpu.Memory[pc]; opcode {
	case 0x31:
		cpu.Ox31()
	case 0xaf:
		cpu.Oxaf()
	case 0x21:
		cpu.Ox21()
	case 0x32:
		cpu.Ox32()
	default:
		cpu.PC += 1
	}

}

func (cpu *CPU) RunBootSequence(){
	for cpu.PC < 256 {
		cpu.FetchDecodeExecute()
	}
}

