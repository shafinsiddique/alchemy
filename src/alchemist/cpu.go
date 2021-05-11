package main

type CPU struct {
	Registers      *Registers
	BootRomMemory  []byte
	Memory         []byte
	PC             uint16
	SP             uint16
	MMU            *MMU
	IME            bool
	Halted         bool
	Debug *bool
	Timer int
	DivideTimer int
	currentTimer int
}

func NewCPU(mmu *MMU) *CPU {
	return &CPU{MMU: mmu, Registers: InitializeRegisters(mmu.Memory)}
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

func (cpu *CPU) storePcAndJump(addr uint16) {
	cpu.PUSH_16(cpu.PC)
	cpu.PC = addr
}

func (cpu *CPU) FetchAndIncrement() byte {
	pc := &cpu.PC
	item := cpu.MMU.Read(*pc)
	*pc += 1
	return item
}

func (cpu *CPU) incrementRegister8Bit(register *EightBitRegister) int {
	initial := register.Get()
	register.Increment()
	cpu.CheckAndSetZeroFlag(register.Get()) // problematic; TODO: fix.
	cpu.CheckAndSetHCFlag(initial, 1, false)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	return 4
}

func (cpu *CPU) decrementRegister(register *EightBitRegister) int {
	initial := register.Get()
	register.Decrement()
	if initial == 1 {
		cpu.SetZeroFlag()
	} else {
		cpu.ClearZeroFlag()
	}
	cpu.CheckAndSetHCFlag(initial, 1, true)
	cpu.SetNegativeFlag()
	return 4
}

func (cpu *CPU) Execute(opcode byte) int {
	//fmt.Println(fmt.Sprintf("Executing Instruction: 0x%x At PC %x and AF %x", opcode, cpu.PC-1, cpu.Registers.AF.Get()))
	//fmt.Println(fmt.Sprintf("Executing Instruction 0x %x At PC %d", opcode, cpu.PC-1))
	cycles := 4 // fetch and increment is a 4.
	switch opcode {
	case 0x31:
		cycles = cpu.LD_SP_D16()
	case 0xaf:
		cycles = cpu.XOR_A()
	case 0x21:
		cycles = cpu.LD_HL_D16()
	case 0x32:
		cycles = cpu.LD_LOC_HL_A_DEC()
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
		cycles = cpu.RLA()
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
	case 0xbe:
		cycles = cpu.CP_LOC_HL()
	case 0x7d:
		cycles = cpu.LD_A_L()
	case 0x78:
		cycles = cpu.LD_A_B()
	case 0x86:
		cycles = cpu.ADD_A_LOC_HL()
	case 0x0:
		cycles = cpu.NOP()
	case 0x1:
		cycles = cpu.LD_BC_D16()
	case 0x2:
		cycles = cpu.LD_LOC_BC_A()
	case 0x3:
		cycles = cpu.INC_BC()
	case 0x7:
		cycles = cpu.RLCA()
	case 0x8:
		cycles = cpu.LD_LOC_A16_SP()
	case 0x9:
		cycles = cpu.ADD_HL_BC()
	case 0xA:
		cycles = cpu.LD_A_LOC_BC()
	case 0xB:
		cycles = cpu.DEC_BC()
	case 0xf:
		cycles = cpu.RRCA()
	case 0x10:
		cycles = cpu.STOP()
	case 0x12:
		cycles = cpu.LD_LOC_DE_A()
	case 0x14:
		cycles = cpu.INC_D()
	case 0x19:
		cycles = cpu.ADD_HL_DE()
	case 0x1B:
		cycles = cpu.DEC_DE()
	case 0x1C:
		cycles = cpu.INC_E()
	case 0x1f:
		cycles = cpu.RRA()
	case 0x25:
		cycles = cpu.DEC_H()
	case 0x26:
		cycles = cpu.LD_H_D8()
	case 0x27:
		cycles = cpu.DAA()
	case 0x29:
		cycles = cpu.ADD_HL_HL()
	case 0x2a:
		cycles = cpu.LD_A_LOC_HL_INC()
	case 0x2b:
		cycles = cpu.DEC_HL()
	case 0x2c:
		cycles = cpu.INC_L()
	case 0x2d:
		cycles = cpu.DEC_L()
	case 0x2f:
		cycles = cpu.CPL()
	case 0x30:
		cycles = cpu.JR_NC_S8()
	case 0x33:
		cycles = cpu.INC_SP()
	case 0x34:
		cycles = cpu.INC_LOC_HL()
	case 0x35:
		cycles = cpu.DEC_LOC_HL()
	case 0x36:
		cycles = cpu.LD_LOC_HL_D8()
	case 0x37:
		cycles = cpu.SCF()
	case 0x38:
		cycles = cpu.JR_C_S8()
	case 0x39:
		cycles = cpu.ADD_HL_SP()
	case 0x3a:
		cycles = cpu.LD_A_LOC_HL_DEC()
	case 0x3b:
		cycles = cpu.DEC_SP()
	case 0x3c:
		cycles = cpu.INC_A()
	case 0x3f:
		cycles = cpu.CCF()
	case 0x40:
		cycles = cpu.LD_B_B()
	case 0x41:
		cycles = cpu.LD_B_C()
	case 0x42:
		cycles = cpu.LD_B_D()
	case 0x43:
		cycles = cpu.LD_B_E()
	case 0x44:
		cycles = cpu.LD_B_H()
	case 0x45:
		cycles = cpu.LD_B_L()
	case 0x46:
		cycles = cpu.LD_B_LOC_HL()
	case 0x47:
		cycles = cpu.LD_B_A()
	case 0x48:
		cycles = cpu.LD_C_B()
	case 0x49:
		cycles = cpu.LD_C_C()
	case 0x4a:
		cycles = cpu.LD_C_D()
	case 0x4b:
		cycles = cpu.LD_C_E()
	case 0x4c:
		cycles = cpu.LD_C_H()
	case 0x4d:
		cycles = cpu.LD_C_L()
	case 0x4e:
		cycles = cpu.LD_C_LOC_HL()
	case 0x50:
		cycles = cpu.LD_D_B()
	case 0x51:
		cycles = cpu.LD_D_C()
	case 0x52:
		cycles = cpu.LD_D_D()
	case 0x53:
		cycles = cpu.LD_D_E()
	case 0x54:
		cycles = cpu.LD_D_H()
	case 0x55:
		cycles = cpu.LD_H_L()
	case 0x56:
		cycles = cpu.LD_D_LOC_HL()
	case 0x58:
		cycles = cpu.LD_E_B()
	case 0x59:
		cycles = cpu.LD_E_C()
	case 0x5a:
		cycles = cpu.LD_E_D()
	case 0x5b:
		cycles = cpu.LD_E_E()
	case 0x5c:
		cycles = cpu.LD_E_H()
	case 0x5d:
		cycles = cpu.LD_E_L()
	case 0x5e:
		cycles = cpu.LD_E_LOC_HL()
	case 0x5f:
		cycles = cpu.LD_E_A()
	case 0x60:
		cycles = cpu.LD_H_B()
	case 0x61:
		cycles = cpu.LD_H_C()
	case 0x62:
		cycles = cpu.LD_H_D()
	case 0x63:
		cycles = cpu.LD_H_E()
	case 0x64:
		cycles = cpu.LD_H_H()
	case 0x65:
		cycles = cpu.LD_H_L()
	case 0x66:
		cycles = cpu.LD_H_LOC_HL()
	case 0x68:
		cycles = cpu.LD_L_B()
	case 0x69:
		cycles = cpu.LD_L_C()
	case 0x6a:
		cycles = cpu.LD_L_D()
	case 0x6b:
		cycles = cpu.LD_L_E()
	case 0x6c:
		cycles = cpu.LD_L_H()
	case 0x6d:
		cycles = cpu.LD_L_L()
	case 0x6e:
		cycles = cpu.LD_L_LOC_HL()
	case 0x6f:
		cycles = cpu.LD_L_A()
	case 0x70:
		cycles = cpu.LD_LOC_HL_B()
	case 0x71:
		cycles = cpu.LD_LOC_HL_C()
	case 0x72:
		cycles = cpu.LD_LOC_HL_D()
	case 0x73:
		cycles = cpu.LD_LOC_HL_E()
	case 0x74:
		cycles = cpu.LD_LOC_HL_H()
	case 0x75:
		cycles = cpu.LD_LOC_HL_L()
	case 0x76:
		cycles = cpu.HALT()
	case 0x79:
		cycles = cpu.LD_A_C()
	case 0x7a:
		cycles = cpu.LD_A_D()
	case 0x7e:
		cycles = cpu.LD_A_LOC_HL()
	case 0x7f:
		cycles = cpu.LD_A_A()
	case 0x80:
		cycles = cpu.ADD_A_B()
	case 0x81:
		cycles = cpu.ADD_A_C()
	case 0x82:
		cycles = cpu.ADD_A_D()
	case 0x83:
		cycles = cpu.ADD_A_E()
	case 0x84:
		cycles = cpu.ADD_A_H()
	case 0x85:
		cycles = cpu.ADD_A_L()
	case 0x87:
		cycles = cpu.ADD_A_A()
	case 0x88:
		cycles = cpu.ADC_A_B()
	case 0x89:
		cycles = cpu.ADC_A_C()
	case 0x8a:
		cycles = cpu.ADC_A_D()
	case 0x8b:
		cycles = cpu.ADC_A_E()
	case 0x8c:
		cycles = cpu.ADC_A_H()
	case 0x8d:
		cycles = cpu.ADC_A_L()
	case 0x8e:
		cycles = cpu.ADC_A_LOC_HL()
	case 0x8f:
		cycles = cpu.ADC_A_A()
	case 0x91:
		cycles = cpu.SUB_C()
	case 0x92:
		cycles = cpu.SUB_D()
	case 0x93:
		cycles = cpu.SUB_E()
	case 0x94:
		cycles = cpu.SUB_H()
	case 0x95:
		cycles = cpu.SUB_L()
	case 0x96:
		cycles = cpu.SUB_LOC_HL()
	case 0x97:
		cycles = cpu.SUB_A()
	case 0x98:
		cycles = cpu.SBC_A_B()
	case 0x99:
		cycles = cpu.SBC_A_C()
	case 0x9a:
		cycles = cpu.SBC_A_D()
	case 0x9b:
		cycles = cpu.SBC_A_E()
	case 0x9c:
		cycles = cpu.SBC_A_H()
	case 0x9d:
		cycles = cpu.SBC_A_L()
	case 0x9e:
		cycles = cpu.SBC_A_LOC_HL()
	case 0x9f:
		cycles = cpu.SBC_A_A()
	case 0xa0:
		cycles = cpu.AND_B()
	case 0xa1:
		cycles = cpu.AND_C()
	case 0xa2:
		cycles = cpu.AND_D()
	case 0xa3:
		cycles = cpu.AND_E()
	case 0xa4:
		cycles = cpu.AND_H()
	case 0xa5:
		cycles = cpu.AND_L()
	case 0xa6:
		cycles = cpu.AND_LOC_HL()
	case 0xa7:
		cycles = cpu.AND_A()
	case 0xa8:
		cycles = cpu.XOR_B()
	case 0xa9:
		cycles = cpu.XOR_C()
	case 0xaa:
		cycles = cpu.XOR_D()
	case 0xab:
		cycles = cpu.XOR_E()
	case 0xac:
		cycles = cpu.XOR_H()
	case 0xad:
		cycles = cpu.XOR_L()
	case 0xae:
		cycles = cpu.XOR_LOC_HL()
	case 0xb0:
		cycles = cpu.OR_B()
	case 0xb1:
		cycles = cpu.OR_C()
	case 0xb2:
		cycles = cpu.OR_D()
	case 0xb3:
		cycles = cpu.OR_E()
	case 0xb4:
		cycles = cpu.OR_H()
	case 0xb5:
		cycles = cpu.OR_L()
	case 0xb6:
		cycles = cpu.OR_LOC_HL()
	case 0xb7:
		cycles = cpu.OR_A()
	case 0xb8:
		cycles = cpu.CP_B()
	case 0xb9:
		cycles = cpu.CP_C()
	case 0xba:
		cycles = cpu.CP_D()
	case 0xbb:
		cycles = cpu.CP_E()
	case 0xbc:
		cycles = cpu.CP_H()
	case 0xbd:
		cycles = cpu.CP_L()
	case 0xbf:
		cycles = cpu.CP_A()
	case 0xc0:
		cycles = cpu.RET_NZ()
	case 0xc2:
		cycles = cpu.JP_NZ_A16()
	case 0xc3:
		cycles = cpu.JP_A16()
	case 0xc4:
		cycles = cpu.CALL_NZ_A16()
	case 0xc6:
		cycles = cpu.ADD_A_D8()
	case 0xc7:
		cycles = cpu.RST_0()
	case 0xc8:
		cycles = cpu.RET_Z()
	case 0xca:
		cycles = cpu.JP_Z_A16()
	case 0xcc:
		cycles = cpu.CALL_Z_A16()
	case 0xce:
		cycles = cpu.ADC_A_D8()
	case 0xcf:
		cycles = cpu.RST_1()
	case 0xd0:
		cycles = cpu.RET_NC()
	case 0xd1:
		cycles = cpu.POP_DE()
	case 0xd2:
		cycles = cpu.JP_NC_A16()
	case 0xd4:
		cycles = cpu.CALL_NC_A16()
	case 0xd5:
		cycles = cpu.PUSH_DE()
	case 0xd6:
		cycles = cpu.SUB_D8()
	case 0xd7:
		cycles = cpu.RST_2()
	case 0xd8:
		cycles = cpu.RET_C()
	case 0xd9:
		cycles = cpu.RETI()
	case 0xda:
		cycles = cpu.JP_C_A16()
	case 0xdc:
		cycles = cpu.CALL_C_A16()
	case 0xde:
		cycles = cpu.SBC_A_D8()
	case 0xdf:
		cycles = cpu.RST_3()
	case 0xe1:
		cycles = cpu.POP_HL()
	case 0xe5:
		cycles = cpu.PUSH_HL()
	case 0xe6:
		cycles = cpu.AND_D8()
	case 0xe7:
		cycles = cpu.RST_4()
	case 0xe8:
		cycles = cpu.ADD_SP_S8()
	case 0xe9:
		cycles = cpu.JP_HL()
	case 0xee:
		cycles = cpu.XOR_D8()
	case 0xef:
		cycles = cpu.RST_5()
	case 0xf1:
		cycles = cpu.POP_AF()
	case 0xf2:
		cycles = cpu.LD_A_C()
	case 0xf3:
		cycles = cpu.DI()
	case 0xf5:
		cycles = cpu.PUSH_AF()
	case 0xf6:
		cycles = cpu.OR_D8()
	case 0xf7:
		cycles = cpu.RST_6()
	case 0xf8:
		cycles = cpu.LD_HL_SP_S8()
	case 0xf9:
		cycles = cpu.LD_SP_HL()
	case 0xfa:
		cycles = cpu.LD_A_LOC_A16()
	case 0xfb:
		cycles = cpu.EI()
	case 0xff:
		cycles = cpu.RST_7()

	default:
		//hex := fmt.Sprintf("0x%x %d", opcode, cpu.PC-1)
		//fmt.Println(hex)
	}
	return cycles
}
