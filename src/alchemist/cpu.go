package main

import "fmt"

type CPU struct {
	Registers *Registers
	Memory []byte
	PC uint16
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

func (cpu *CPU) PopFromStack()byte {
	sp := &cpu.SP
	item := cpu.Memory[*sp]
	*sp += 1
	return item
}

func (cpu *CPU) FetchAndIncrement() byte {
	pc := &cpu.PC
	item := cpu.Memory[*pc]
	*pc += 1
	return item
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
	opcode := cpu.FetchAndIncrement()
	fmt.Println(fmt.Sprintf("Executing Instruction %x At PC %d",opcode, cpu.PC - 1))
	switch opcode {
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
	case 0x0c:
		cpu.INC_C()
	case 0x77:
		cpu.LD_LOC_HL_A()
	case 0xe0:
		cpu.LD_LOC_A8_A()
	case 0x11:
		cpu.LD_DE_D16()
	case 0x1a:
		cpu.LD_A_LOC_DE()
	case 0x4f:
		cpu.LD_C_A()
	case 0x06:
		cpu.LD_B_D8()
	case 0xc5:
		cpu.PUSH_BC()
	case 0x17:
		cpu.RLA()
	case 0xc1:
		cpu.POP_BC()
	case 0x5:
		cpu.DEC_B()
	case 0x22:
		cpu.LD_LOC_HL_A_INC()
	case 0x23:
		cpu.INC_HL()
	case 0xc9:
		cpu.RET()
	case 0x13:
		cpu.INC_DE()
	case 0x7b:
		cpu.LD_A_E()
	case 0xfe:
		cpu.CP_D8()
	case 0xea:
		cpu.LD_LOC_A16_A()
	case 0x3d:
		cpu.DEC_A()
	case 0x28:
		cpu.JR_Z_S8()
	case 0x67:
		cpu.LD_H_A()
	case 0x57:
		cpu.LD_D_A()
	case 0x4:
		cpu.INC_B()
	case 0x1e:
		cpu.LD_E_D8()
	case 0xf0:
		cpu.LD_A_LOC_A8()
	case 0xd:
		cpu.DEC_C()
	case 0x1d:
		cpu.DEC_E()
	case 0x15:
		cpu.DEC_D()
	case 0x16:
		cpu.LD_D_D8()
	default:
		hex := fmt.Sprintf("0x%x %d", opcode, cpu.PC-1)
		fmt.Println(hex)
	}
}

func (cpu *CPU) RunBootSequence(){
	for cpu.PC < 256 {
		cpu.FetchDecodeExecute()
	}
}



