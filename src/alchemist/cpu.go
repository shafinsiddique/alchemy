package main

import (
	"fmt"
)

type CPU struct {
	Registers     *Registers
	BootRomMemory []byte
	Memory        []byte
	PC            uint16
	SP            uint16
	MMU           *MMU
}

func (cpu *CPU) PushToStack(item byte) {
	sp := &cpu.SP
	*sp -= 1
	cpu.MMU.Write(*sp, item)
}

func (cpu *CPU) PopFromStack() byte {
	sp := &cpu.SP
	item := cpu.MMU.Read(*sp)
	*sp += 1
	return item
}

func (cpu *CPU) FetchAndIncrement() byte {
	pc := &cpu.PC
	item := cpu.MMU.Read(*pc)
	*pc += 1
	return item
}

func (cpu *CPU) IncrementRegister8Bit(register *EightBitRegister) int {
	initial := register.Get()
	cycles := register.Increment()
	cpu.CheckAndSetZeroFlag(register.Get())
	cpu.CheckAndSetHCFlag(initial, 1, false)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	return cycles
}

func (cpu *CPU) DecrementRegister8Bit(register *EightBitRegister) int {
	initial := register.Get()
	if initial == 0 {
		cpu.ClearHCFlag()
		cpu.ClearZeroFlag()
	} else {
		register.Decrement()

		if initial == 1 {
			cpu.SetZeroFlag()
		} else {
			cpu.ClearZeroFlag()
		}

		cpu.CheckAndSetHCFlag(initial, 1, true)
	}
	cpu.Registers.F.SetBit(1, NEGATIVE_FLAG)
	return 4
}

func (cpu *CPU) FetchDecodeExecute() int {
	opcode := cpu.FetchAndIncrement()
	fmt.Println(fmt.Sprintf("Executing Instruction 0x %x At PC %d", opcode, cpu.PC-1))
	//logger.Info(fmt.Sprintf("Executing Instruction 0x %x At PC %x", opcode, cpu.PC-1))
	cycles := 4 // fetch and increment is a 4.
	switch opcode {
	case 0x31:
		cycles = cpu.LD_SP_D16()
	case 0xaf:
		cycles = cpu.XOR_A()
	case 0x21:
		cycles = cpu.LD_HL_D16()
	case 0x32:
		cycles = cpu.LD_HL_A_DEC()
	case 0xcb:
		cycles = cpu.Oxcb()
	case 0x20:
		cycles = cpu.JR_NZ_S8() // s8 stands for signed 8 bit.
	case 0x0E:
		cycles = cpu.LD_C_D8()
	case 0x3e:
		cycles = cpu.LD_A_D8()
	case 0xe2:
		cycles = cpu.LD_LOC_C_A()
	case 0xcd:
		cycles = cpu.CALL_A16()
	case 0x0c:
		cycles = cpu.INC_C()
	case 0x77:
		cycles = cpu.LD_LOC_HL_A()
	case 0xe0:
		cycles = cpu.LD_LOC_A8_A()
	case 0x11:
		cycles = cpu.LD_DE_D16()
	case 0x1a:
		cycles = cpu.LD_A_LOC_DE()
	case 0x4f:
		cycles = cpu.LD_C_A()
	case 0x06:
		cycles = cpu.LD_B_D8()
	case 0xc5:
		cycles = cpu.PUSH_BC()
	case 0x17:
		cycles = cpu.RL_A()
	case 0xc1:
		cycles = cpu.POP_BC()
	case 0x5:
		cycles = cpu.DEC_B()
	case 0x22:
		cycles = cpu.LD_LOC_HL_A_INC()
	case 0x23:
		cycles = cpu.INC_HL()
	case 0xc9:
		cycles = cpu.RET()
	case 0x13:
		cycles = cpu.INC_DE()
	case 0x7b:
		cycles = cpu.LD_A_E()
	case 0xfe:
		cycles = cpu.CP_D8()
	case 0xea:
		cycles = cpu.LD_LOC_A16_A()
	case 0x3d:
		cycles = cpu.DEC_A()
	case 0x28:
		cycles = cpu.JR_Z_S8()
	case 0x67:
		cycles = cpu.LD_H_A()
	case 0x57:
		cycles = cpu.LD_D_A()
	case 0x4:
		cycles = cpu.INC_B()
	case 0x1e:
		cycles = cpu.LD_E_D8()
	case 0xf0:
		cycles = cpu.LD_A_LOC_A8()
	case 0xd:
		cycles = cpu.DEC_C()
	case 0x1d:
		cycles = cpu.DEC_E()
	case 0x15:
		cycles = cpu.DEC_D()
	case 0x16:
		cycles = cpu.LD_D_D8()
	case 0x24:
		cycles = cpu.INC_H()
	case 0x7c:
		cycles = cpu.LD_A_H()
	case 0x90:
		cycles = cpu.SUB_B()
	case 0x18:
		cycles = cpu.JR_S8()
	case 0x2e:
		cycles = cpu.LD_L_D8()
	default:
		hex := fmt.Sprintf("0x%x %d", opcode, cpu.PC-1)
		fmt.Println(hex)
	}
	return cycles
}


