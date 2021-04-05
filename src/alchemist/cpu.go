package main

import "fmt"

type CPU struct {
	Registers *Registers
	Memory []byte
	PC int16
	SP uint16
	incremented bool
}

func NewCPU() *CPU {
	return &CPU{Registers: InitializeRegisters(), Memory: make([]byte,0x10000)}
}

func (cpu *CPU) LoadBootRom(bootrom []byte) {
	for i := 0; i < len(bootrom) ; i++ {
		cpu.Memory[i] = bootrom[i]
	}
}

func (cpu *CPU) PushToStack(item byte) {
	sp := &cpu.SP
	*sp -= 1
	cpu.Memory[*sp] = item
}

func (cpu *CPU) IncrementPC()  {
	cpu.PC += 1
}

func (cpu *CPU) DecrementPC(){
	cpu.PC -=1
}

func (cpu *CPU) GetElementAtPC() byte {
	return cpu.Memory[cpu.PC]
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
		cpu.LD_HL_A_DEC()
	case 0xcb:
		cpu.Oxcb()
	case 0x20:
		cpu.JR_NZ_S8() // s8 stands for signed 8 bit.
	case 0x0E:
		cpu.LD_C_D8()
	case 0x3e:
		cpu.LD_A_D8()
	case 0xe2:
		cpu.LD_LOC_C_A()
	case 0xcd:
		cpu.CALL_A16()
	default:
		hex := fmt.Sprintf("0x%x %d", opcode, *pc)
		fmt.Println(hex)
	}
	*pc += 1

}

func (cpu *CPU) RunBootSequence(){
	for cpu.PC < 256 {
		cpu.FetchDecodeExecute()
	}
}

