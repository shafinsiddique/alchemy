package main

type cbFun func(byte) byte

func (cpu *CPU) Oxcb() int {
	cycles := 8
	switch opcode := cpu.FetchAndIncrement(); opcode {
	case 0xc0:
		cycles = cpu.SET_0_B()
	case 0xc1:
		cycles = cpu.SET_0_C()
	case 0xc2:
		cycles = cpu.SET_0_D()
	case 0xc3:
		cycles = cpu.SET_0_E()
	case 0xc4:
		cycles = cpu.SET_0_H()
	case 0xc5:
		cycles = cpu.SET_0_L()
	case 0xc6:
		cycles = cpu.SET_0_HL()
	case 0xc7:
		cycles = cpu.SET_0_A()
	case 0xc8:
		cycles = cpu.SET_1_B()
	case 0xc9:
		cycles = cpu.SET_1_C()
	case 0xca:
		cycles = cpu.SET_1_D()
	case 0xcb:
		cycles = cpu.SET_1_E()
	case 0xcc:
		cycles = cpu.SET_1_H()
	case 0xcd:
		cycles = cpu.SET_1_L()
	case 0xce:
		cycles = cpu.SET_1_HL()
	case 0xcf:
		cycles = cpu.SET_1_A()
	case 0xd0:
		cycles = cpu.SET_2_B()
	case 0xd1:
		cycles = cpu.SET_2_C()
	case 0xd2:
		cycles = cpu.SET_2_D()
	case 0xd3:
		cycles = cpu.SET_2_E()
	case 0xd4:
		cycles = cpu.SET_2_H()
	case 0xd5:
		cycles = cpu.SET_2_L()
	case 0xd6:
		cycles = cpu.SET_2_HL()
	case 0xd7:
		cycles = cpu.SET_2_A()
	case 0xd8:
		cycles = cpu.SET_3_B()
	case 0xd9:
		cycles = cpu.SET_3_C()
	case 0xda:
		cycles = cpu.SET_3_D()
	case 0xdb:
		cycles = cpu.SET_3_E()
	case 0xdc:
		cycles = cpu.SET_3_H()
	case 0xdd:
		cycles = cpu.SET_3_L()
	case 0xde:
		cycles = cpu.SET_3_HL()
	case 0xdf:
		cycles = cpu.SET_3_A()
	case 0xe0:
		cycles = cpu.SET_4_B()
	case 0xe1:
		cycles = cpu.SET_4_C()
	case 0xe2:
		cycles = cpu.SET_4_D()
	case 0xe3:
		cycles = cpu.SET_4_E()
	case 0xe4:
		cycles = cpu.SET_4_H()
	case 0xe5:
		cycles = cpu.SET_4_L()
	case 0xe6:
		cycles = cpu.SET_4_HL()
	case 0xe7:
		cycles = cpu.SET_4_A()
	case 0xe8:
		cycles = cpu.SET_5_B()
	case 0xe9:
		cycles = cpu.SET_5_C()
	case 0xea:
		cycles = cpu.SET_5_D()
	case 0xeb:
		cycles = cpu.SET_5_E()
	case 0xec:
		cycles = cpu.SET_5_H()
	case 0xed:
		cycles = cpu.SET_5_L()
	case 0xee:
		cycles = cpu.SET_5_HL()
	case 0xef:
		cycles = cpu.SET_5_A()
	case 0xf0:
		cycles = cpu.SET_6_B()
	case 0xf1:
		cycles = cpu.SET_6_C()
	case 0xf2:
		cycles = cpu.SET_6_D()
	case 0xf3:
		cycles = cpu.SET_6_E()
	case 0xf4:
		cycles = cpu.SET_6_H()
	case 0xf5:
		cycles = cpu.SET_6_L()
	case 0xf6:
		cycles = cpu.SET_6_HL()
	case 0xf7:
		cycles = cpu.SET_6_A()
	case 0xf8:
		cycles = cpu.SET_7_B()
	case 0xf9:
		cycles = cpu.SET_7_C()
	case 0xfa:
		cycles = cpu.SET_7_D()
	case 0xfb:
		cycles = cpu.SET_7_E()
	case 0xfc:
		cycles = cpu.SET_7_H()
	case 0xfd:
		cycles = cpu.SET_7_L()
	case 0xfe:
		cycles = cpu.SET_7_HL()
	case 0xff:
		cycles = cpu.SET_7_A()
	case 0x80:
		cycles=cpu.RES_0_B()
	case 0x81:
		cycles=cpu.RES_0_C()
	case 0x82:
		cycles=cpu.RES_0_D()
	case 0x83:
		cycles=cpu.RES_0_E()
	case 0x84:
		cycles=cpu.RES_0_H()
	case 0x85:
		cycles=cpu.RES_0_L()
	case 0x86:
		cycles=cpu.RES_0_HL()
	case 0x87:
		cycles=cpu.RES_0_A()
	case 0x88:
		cycles=cpu.RES_1_B()
	case 0x89:
		cycles=cpu.RES_1_C()
	case 0x8a:
		cycles=cpu.RES_1_D()
	case 0x8b:
		cycles=cpu.RES_1_E()
	case 0x8c:
		cycles=cpu.RES_1_H()
	case 0x8d:
		cycles=cpu.RES_1_L()
	case 0x8e:
		cycles=cpu.RES_1_HL()
	case 0x8f:
		cycles=cpu.RES_1_A()
	case 0x90:
		cycles=cpu.RES_2_B()
	case 0x91:
		cycles=cpu.RES_2_C()
	case 0x92:
		cycles=cpu.RES_2_D()
	case 0x93:
		cycles=cpu.RES_2_E()
	case 0x94:
		cycles=cpu.RES_2_H()
	case 0x95:
		cycles=cpu.RES_2_L()
	case 0x96:
		cycles=cpu.RES_2_HL()
	case 0x97:
		cycles=cpu.RES_2_A()
	case 0x98:
		cycles=cpu.RES_3_B()
	case 0x99:
		cycles=cpu.RES_3_C()
	case 0x9a:
		cycles=cpu.RES_3_D()
	case 0x9b:
		cycles=cpu.RES_3_E()
	case 0x9c:
		cycles=cpu.RES_3_H()
	case 0x9d:
		cycles=cpu.RES_3_L()
	case 0x9e:
		cycles=cpu.RES_3_HL()
	case 0x9f:
		cycles=cpu.RES_3_A()
	case 0xa0:
		cycles=cpu.RES_4_B()
	case 0xa1:
		cycles=cpu.RES_4_C()
	case 0xa2:
		cycles=cpu.RES_4_D()
	case 0xa3:
		cycles=cpu.RES_4_E()
	case 0xa4:
		cycles=cpu.RES_4_H()
	case 0xa5:
		cycles=cpu.RES_4_L()
	case 0xa6:
		cycles=cpu.RES_4_HL()
	case 0xa7:
		cycles=cpu.RES_4_A()
	case 0xa8:
		cycles=cpu.RES_5_B()
	case 0xa9:
		cycles=cpu.RES_5_C()
	case 0xaa:
		cycles=cpu.RES_5_D()
	case 0xab:
		cycles=cpu.RES_5_E()
	case 0xac:
		cycles=cpu.RES_5_H()
	case 0xad:
		cycles=cpu.RES_5_L()
	case 0xae:
		cycles=cpu.RES_5_HL()
	case 0xaf:
		cycles=cpu.RES_5_A()
	case 0xb0:
		cycles=cpu.RES_6_B()
	case 0xb1:
		cycles=cpu.RES_6_C()
	case 0xb2:
		cycles=cpu.RES_6_D()
	case 0xb3:
		cycles=cpu.RES_6_E()
	case 0xb4:
		cycles=cpu.RES_6_H()
	case 0xb5:
		cycles=cpu.RES_6_L()
	case 0xb6:
		cycles=cpu.RES_6_HL()
	case 0xb7:
		cycles=cpu.RES_6_A()
	case 0xb8:
		cycles=cpu.RES_7_B()
	case 0xb9:
		cycles=cpu.RES_7_C()
	case 0xba:
		cycles=cpu.RES_7_D()
	case 0xbb:
		cycles=cpu.RES_7_E()
	case 0xbc:
		cycles=cpu.RES_7_H()
	case 0xbd:
		cycles=cpu.RES_7_L()
	case 0xbe:
		cycles=cpu.RES_7_HL()
	case 0xbf:
		cycles=cpu.RES_7_A()
	case 0x40:
		cycles=cpu.BIT_0_B()
	case 0x41:
		cycles=cpu.BIT_0_C()
	case 0x42:
		cycles=cpu.BIT_0_D()
	case 0x43:
		cycles=cpu.BIT_0_E()
	case 0x44:
		cycles=cpu.BIT_0_H()
	case 0x45:
		cycles=cpu.BIT_0_L()
	case 0x46:
		cycles=cpu.BIT_0_HL()
	case 0x47:
		cycles=cpu.BIT_0_A()
	case 0x48:
		cycles=cpu.BIT_1_B()
	case 0x49:
		cycles=cpu.BIT_1_C()
	case 0x4a:
		cycles=cpu.BIT_1_D()
	case 0x4b:
		cycles=cpu.BIT_1_E()
	case 0x4c:
		cycles=cpu.BIT_1_H()
	case 0x4d:
		cycles=cpu.BIT_1_L()
	case 0x4e:
		cycles=cpu.BIT_1_HL()
	case 0x4f:
		cycles=cpu.BIT_1_A()
	case 0x50:
		cycles=cpu.BIT_2_B()
	case 0x51:
		cycles=cpu.BIT_2_C()
	case 0x52:
		cycles=cpu.BIT_2_D()
	case 0x53:
		cycles=cpu.BIT_2_E()
	case 0x54:
		cycles=cpu.BIT_2_H()
	case 0x55:
		cycles=cpu.BIT_2_L()
	case 0x56:
		cycles=cpu.BIT_2_HL()
	case 0x57:
		cycles=cpu.BIT_2_A()
	case 0x58:
		cycles=cpu.BIT_3_B()
	case 0x59:
		cycles=cpu.BIT_3_C()
	case 0x5a:
		cycles=cpu.BIT_3_D()
	case 0x5b:
		cycles=cpu.BIT_3_E()
	case 0x5c:
		cycles=cpu.BIT_3_H()
	case 0x5d:
		cycles=cpu.BIT_3_L()
	case 0x5e:
		cycles=cpu.BIT_3_HL()
	case 0x5f:
		cycles=cpu.BIT_3_A()
	case 0x60:
		cycles=cpu.BIT_4_B()
	case 0x61:
		cycles=cpu.BIT_4_C()
	case 0x62:
		cycles=cpu.BIT_4_D()
	case 0x63:
		cycles=cpu.BIT_4_E()
	case 0x64:
		cycles=cpu.BIT_4_H()
	case 0x65:
		cycles=cpu.BIT_4_L()
	case 0x66:
		cycles=cpu.BIT_4_HL()
	case 0x67:
		cycles=cpu.BIT_4_A()
	case 0x68:
		cycles=cpu.BIT_5_B()
	case 0x69:
		cycles=cpu.BIT_5_C()
	case 0x6a:
		cycles=cpu.BIT_5_D()
	case 0x6b:
		cycles=cpu.BIT_5_E()
	case 0x6c:
		cycles=cpu.BIT_5_H()
	case 0x6d:
		cycles=cpu.BIT_5_L()
	case 0x6e:
		cycles=cpu.BIT_5_HL()
	case 0x6f:
		cycles=cpu.BIT_5_A()
	case 0x70:
		cycles=cpu.BIT_6_B()
	case 0x71:
		cycles=cpu.BIT_6_C()
	case 0x72:
		cycles=cpu.BIT_6_D()
	case 0x73:
		cycles=cpu.BIT_6_E()
	case 0x74:
		cycles=cpu.BIT_6_H()
	case 0x75:
		cycles=cpu.BIT_6_L()
	case 0x76:
		cycles=cpu.BIT_6_HL()
	case 0x77:
		cycles=cpu.BIT_6_A()
	case 0x78:
		cycles=cpu.BIT_7_B()
	case 0x79:
		cycles=cpu.BIT_7_C()
	case 0x7a:
		cycles=cpu.BIT_7_D()
	case 0x7b:
		cycles=cpu.BIT_7_E()
	case 0x7c:
		cycles=cpu.BIT_7_H()
	case 0x7d:
		cycles=cpu.BIT_7_L()
	case 0x7e:
		cycles=cpu.BIT_7_HL()
	case 0x7f:
		cycles=cpu.BIT_7_A()
	case 0x30:
		cycles=cpu.SWAP_B()
	case 0x31:
		cycles=cpu.SWAP_C()
	case 0x32:
		cycles=cpu.SWAP_D()
	case 0x33:
		cycles=cpu.SWAP_E()
	case 0x34:
		cycles=cpu.SWAP_H()
	case 0x35:
		cycles=cpu.SWAP_L()
	case 0x36:
		cycles=cpu.SWAP_HL()
	case 0x37:
		cycles=cpu.SWAP_A()
	case 0x38:
		cycles=cpu.SRL_B()
	case 0x39:
		cycles=cpu.SRL_C()
	case 0x3a:
		cycles=cpu.SRL_D()
	case 0x3b:
		cycles=cpu.SRL_E()
	case 0x3c:
		cycles=cpu.SRL_H()
	case 0x3d:
		cycles=cpu.SRL_L()
	case 0x3e:
		cycles=cpu.SRL_HL()
	case 0x3f:
		cycles=cpu.SRL_A()
	case 0x20:
		cycles=cpu.SLA_B()
	case 0x21:
		cycles=cpu.SLA_C()
	case 0x22:
		cycles=cpu.SLA_D()
	case 0x23:
		cycles=cpu.SLA_E()
	case 0x24:
		cycles=cpu.SLA_H()
	case 0x25:
		cycles=cpu.SLA_L()
	case 0x26:
		cycles=cpu.SLA_HL()
	case 0x27:
		cycles=cpu.SLA_A()
	case 0x28:
		cycles=cpu.SRA_B()
	case 0x29:
		cycles=cpu.SRA_C()
	case 0x2a:
		cycles=cpu.SRA_D()
	case 0x2b:
		cycles=cpu.SRA_E()
	case 0x2c:
		cycles=cpu.SRA_H()
	case 0x2d:
		cycles=cpu.SRA_L()
	case 0x2e:
		cycles=cpu.SRA_HL()
	case 0x2f:
		cycles=cpu.SRA_A()
	case 0x10:
		cycles=cpu.RL_B()
	case 0x11:
		cycles=cpu.RL_C()
	case 0x12:
		cycles=cpu.RL_D()
	case 0x13:
		cycles=cpu.RL_E()
	case 0x14:
		cycles=cpu.RL_H()
	case 0x15:
		cycles=cpu.RL_L()
	case 0x16:
		cycles=cpu.RL_HL()
	case 0x17:
		cycles=cpu.RL_A()
	case 0x18:
		cycles=cpu.RR_B()
	case 0x19:
		cycles=cpu.RR_C()
	case 0x1a:
		cycles=cpu.RR_D()
	case 0x1b:
		cycles=cpu.RR_E()
	case 0x1c:
		cycles=cpu.RR_H()
	case 0x1d:
		cycles=cpu.RR_L()
	case 0x1e:
		cycles=cpu.RR_HL()
	case 0x1f:
		cycles=cpu.RR_A()
	case 0x0:
		cycles=cpu.RLC_B()
	case 0x1:
		cycles=cpu.RLC_C()
	case 0x2:
		cycles=cpu.RLC_D()
	case 0x3:
		cycles=cpu.RLC_E()
	case 0x4:
		cycles=cpu.RLC_H()
	case 0x5:
		cycles=cpu.RLC_L()
	case 0x6:
		cycles=cpu.RLC_HL()
	case 0x7:
		cycles=cpu.RLC_A()
	case 0x8:
		cycles=cpu.RRC_B()
	case 0x9:
		cycles=cpu.RRC_C()
	case 0xa:
		cycles=cpu.RRC_D()
	case 0xb:
		cycles=cpu.RRC_E()
	case 0xc:
		cycles=cpu.RRC_H()
	case 0xd:
		cycles=cpu.RRC_L()
	case 0xe:
		cycles=cpu.RRC_HL()
	case 0xf:
		cycles=cpu.RRC_A()

	}

	return cycles
}

func (cpu *CPU) get_rl(value byte) byte {
	//rotate the contents of this value to the leftr through the carry flag.
	//previos contents of cy are in bit 0.
	for i := 0; i <= 7; i++ {
		bit := GetBit(value, i)
		carry := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
		value = SetBit(value, carry, i)
	}

	return value
}

func (cpu *CPU) get_rr(value byte) byte {
	for i := 7; i >= 0; i-- {
		bit := GetBit(value, i)
		carry := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
		value = SetBit(value, carry, i)
	}

	return value
}

func (cpu *CPU) cb_rl8(register *EightBitRegister) int {
	register.Set(cpu.get_rl(register.Get()))
	cpu.CheckAndSetZeroFlag(register.Get())
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 8
}

func (cpu *CPU) cb_rl16(register *SixteenBitRegister) int {
	loc := register.Get()
	rotated := cpu.get_rl(cpu.MMU.Read(loc))
	cpu.MMU.Write(loc, rotated)
	cpu.CheckAndSetZeroFlag(rotated)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 16
}

func (cpu *CPU) cb_rr8(register *EightBitRegister) int {
	register.Set(cpu.get_rr(register.Get()))
	cpu.CheckAndSetZeroFlag(register.Get())
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 8
}

func (cpu *CPU) cb_rr16(register *SixteenBitRegister) int {
	loc := register.Get()
	rotated := cpu.get_rr(cpu.MMU.Read(loc))
	cpu.MMU.Write(loc, rotated)
	cpu.CheckAndSetZeroFlag(rotated)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 16
}

func (cpu *CPU) cb_rlc8(register *EightBitRegister) int {
	cpu.Registers.F.SetBit(register.GetBit(7), CARRY_FLAG)
	return cpu.cb_rl8(register)
}

func (cpu *CPU) cb_rlc16(register *SixteenBitRegister) int {
	loc := register.Get()
	val := cpu.MMU.Read(loc)
	cpu.Registers.F.SetBit(GetBit(val, 7), CARRY_FLAG)
	return cpu.cb_rl16(register)
}

func (cpu *CPU) cb_rrc8(register *EightBitRegister) int {
	cpu.Registers.F.SetBit(register.GetBit(0), CARRY_FLAG)
	return cpu.cb_rr8(register)
}

func (cpu *CPU) cb_rrc16(register *SixteenBitRegister) int {
	loc := register.Get()
	val := cpu.MMU.Read(loc)
	cpu.Registers.F.SetBit(GetBit(val, 0), CARRY_FLAG)
	return cpu.cb_rr16(register)
}

func (cpu *CPU) cb_sla8(register *EightBitRegister) int {
	cpu.Registers.F.SetBit(0, CARRY_FLAG)
	return cpu.cb_rl8(register)
}

func (cpu *CPU) cb_sla16(register *SixteenBitRegister) int {
	cpu.Registers.F.SetBit(0, CARRY_FLAG)
	return cpu.cb_rl16(register)
}

func (cpu *CPU) cb_sra8(register *EightBitRegister) int {
	original := register.GetBit(7)
	cpu.Registers.F.SetBit(original, CARRY_FLAG) // this will get passed to bit 7 ensuring bit 7 stays
	// unchanged,
	return cpu.cb_rr8(register)
}

func (cpu *CPU) cb_sra16(register *SixteenBitRegister) int {
	loc := register.Get()
	val := cpu.MMU.Read(loc)
	original := GetBit(val, 7)
	cpu.Registers.F.SetBit(original, CARRY_FLAG) // this will get passed to bit 7 ensuring bit 7 stays
	// unchanged,
	return cpu.cb_rr16(register)
}

func (cpu *CPU) swap(val byte) byte {
	bits := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bits[i] = GetBit(val, i)
	}

	val = SetBit(val, bits[4], 0) // read as : Set Bit 0 to the bit at index 4.
	val = SetBit(val, bits[5], 1)
	val = SetBit(val, bits[6], 2)
	val = SetBit(val, bits[7], 3)
	val = SetBit(val, bits[0], 4)
	val = SetBit(val, bits[1], 5)
	val = SetBit(val, bits[2], 6)
	val = SetBit(val, bits[3], 7) // read as : Set Bit 7 to the bit at index 3.

	return val
}

func (cpu *CPU) SetRegisterOrLoc(register interface{}, operation cbFun) byte {
	var result byte
	switch register.(type) {
	case *EightBitRegister:
		reg := register.(*EightBitRegister)
		result = operation(reg.Get())
		reg.Set(result)
	default:
		reg := register.(*SixteenBitRegister)
		loc := reg.Get()
		val := cpu.MMU.Read(loc)
		newVal := operation(val)
		result = newVal
		cpu.MMU.Write(loc, newVal)
	}
	return result
}

func (cpu *CPU) cb_swap8(register *EightBitRegister) int {
	cpu.SetRegisterOrLoc(register, cpu.swap)
	return 8
}

func (cpu *CPU) cb_swap16(register *SixteenBitRegister) int {
	cpu.SetRegisterOrLoc(register, cpu.swap)
	return 16
}

func (cpu *CPU) cb_swap(register interface{}) int {
	result := cpu.SetRegisterOrLoc(register, cpu.swap)
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	cpu.ClearOverflowFlag()
	switch register.(type) {
	case *EightBitRegister:
		return 8
	default:
		return 16
	}
}

func (cpu *CPU) cb_srl(register interface{}) int {
	cpu.Registers.F.SetBit(0, CARRY_FLAG)                // ensuring bit 7 is set to 0.
	result := cpu.SetRegisterOrLoc(register, cpu.get_rl) // carry flag will be set here.
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	switch register.(type) {
	case *EightBitRegister:
		return 8
	default:
		return 16
	}
}

func (cpu *CPU) cb_bit(index int, register interface{}) int {
	var value byte
	switch register.(type) {
	case *EightBitRegister:
		reg := register.(*EightBitRegister)
		value = reg.Get()
	default:
		reg := register.(*SixteenBitRegister)
		loc := reg.Get()
		value = cpu.MMU.Read(loc)
	}
	complement := GetBit(value, index) ^ 1
	cpu.Registers.F.SetBit(complement, Z_FLAG)
	cpu.ClearNegativeFlag()
	cpu.SetHCFlag()
	return 8
}

func (cpu *CPU) cb_res_set(index int, bit byte, register interface{}) int {
	cycles := 8
	switch register.(type) {
	case *EightBitRegister:
		reg := register.(*EightBitRegister)
		reg.SetBit(bit, index)
	default:
		reg := register.(*SixteenBitRegister)
		loc := reg.Get()
		cpu.MMU.Write(loc, SetBit(cpu.MMU.Read(loc), bit, index))
		cycles = 16
	}
	return cycles
}

func (cpu *CPU) RES_0_B() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_0_C() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_0_D() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_0_E() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_0_H() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_0_L() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_0_HL() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_0_A() int {
	return cpu.cb_res_set(0, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_1_B() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_1_C() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_1_D() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_1_E() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_1_H() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_1_L() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_1_HL() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_1_A() int {
	return cpu.cb_res_set(1, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_2_B() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_2_C() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_2_D() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_2_E() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_2_H() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_2_L() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_2_HL() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_2_A() int {
	return cpu.cb_res_set(2, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_3_B() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_3_C() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_3_D() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_3_E() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_3_H() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_3_L() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_3_HL() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_3_A() int {
	return cpu.cb_res_set(3, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_4_B() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_4_C() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_4_D() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_4_E() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_4_H() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_4_L() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_4_HL() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_4_A() int {
	return cpu.cb_res_set(4, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_5_B() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_5_C() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_5_D() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_5_E() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_5_H() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_5_L() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_5_HL() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_5_A() int {
	return cpu.cb_res_set(5, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_6_B() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_6_C() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_6_D() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_6_E() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_6_H() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_6_L() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_6_HL() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_6_A() int {
	return cpu.cb_res_set(6, 0, cpu.Registers.A)
}

func (cpu *CPU) RES_7_B() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.B)
}

func (cpu *CPU) RES_7_C() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.C)
}

func (cpu *CPU) RES_7_D() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.D)
}

func (cpu *CPU) RES_7_E() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.E)
}

func (cpu *CPU) RES_7_H() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.H)
}

func (cpu *CPU) RES_7_L() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.L)
}

func (cpu *CPU) RES_7_HL() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.HL)
}

func (cpu *CPU) RES_7_A() int {
	return cpu.cb_res_set(7, 0, cpu.Registers.A)
}

func (cpu *CPU) SET_0_B() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_0_C() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_0_D() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_0_E() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_0_H() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_0_L() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_0_HL() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_0_A() int {
	return cpu.cb_res_set(0, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_1_B() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_1_C() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_1_D() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_1_E() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_1_H() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_1_L() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_1_HL() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_1_A() int {
	return cpu.cb_res_set(1, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_2_B() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_2_C() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_2_D() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_2_E() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_2_H() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_2_L() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_2_HL() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_2_A() int {
	return cpu.cb_res_set(2, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_3_B() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_3_C() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_3_D() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_3_E() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_3_H() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_3_L() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_3_HL() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_3_A() int {
	return cpu.cb_res_set(3, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_4_B() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_4_C() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_4_D() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_4_E() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_4_H() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_4_L() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_4_HL() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_4_A() int {
	return cpu.cb_res_set(4, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_5_B() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_5_C() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_5_D() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_5_E() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_5_H() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_5_L() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_5_HL() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_5_A() int {
	return cpu.cb_res_set(5, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_6_B() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_6_C() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_6_D() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_6_E() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_6_H() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_6_L() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_6_HL() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_6_A() int {
	return cpu.cb_res_set(6, 1, cpu.Registers.A)
}

func (cpu *CPU) SET_7_B() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.B)
}

func (cpu *CPU) SET_7_C() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.C)
}

func (cpu *CPU) SET_7_D() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.D)
}

func (cpu *CPU) SET_7_E() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.E)
}

func (cpu *CPU) SET_7_H() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.H)
}

func (cpu *CPU) SET_7_L() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.L)
}

func (cpu *CPU) SET_7_HL() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_7_A() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.A)
}

func (cpu *CPU) SLA_B() int {
	return cpu.cb_sla8(cpu.Registers.B)
}
func (cpu *CPU) SLA_C() int {
	return cpu.cb_sla8(cpu.Registers.C)
}
func (cpu *CPU) SLA_D() int {
	return cpu.cb_sla8(cpu.Registers.D)
}
func (cpu *CPU) SLA_E() int {
	return cpu.cb_sla8(cpu.Registers.E)
}
func (cpu *CPU) SLA_H() int {
	return cpu.cb_sla8(cpu.Registers.H)
}
func (cpu *CPU) SLA_L() int {
	return cpu.cb_sla8(cpu.Registers.L)
}
func (cpu *CPU) SLA_HL() int {
	return cpu.cb_sla16(cpu.Registers.HL)
}
func (cpu *CPU) SLA_A() int {
	return cpu.cb_sla8(cpu.Registers.A)
}
func (cpu *CPU) SRA_B() int {
	return cpu.cb_sra8(cpu.Registers.B)
}
func (cpu *CPU) SRA_C() int {
	return cpu.cb_sra8(cpu.Registers.C)
}
func (cpu *CPU) SRA_D() int {
	return cpu.cb_sra8(cpu.Registers.D)
}
func (cpu *CPU) SRA_E() int {
	return cpu.cb_sra8(cpu.Registers.E)
}
func (cpu *CPU) SRA_H() int {
	return cpu.cb_sra8(cpu.Registers.H)
}
func (cpu *CPU) SRA_L() int {
	return cpu.cb_sra8(cpu.Registers.L)
}
func (cpu *CPU) SRA_HL() int {
	return cpu.cb_sra16(cpu.Registers.HL)
}
func (cpu *CPU) SRA_A() int {
	return cpu.cb_sra8(cpu.Registers.A)
}
func (cpu *CPU) RL_B() int {
	return cpu.cb_rl8(cpu.Registers.B)
}
func (cpu *CPU) RL_C() int {
	return cpu.cb_rl8(cpu.Registers.C)
}
func (cpu *CPU) RL_D() int {
	return cpu.cb_rl8(cpu.Registers.D)
}
func (cpu *CPU) RL_E() int {
	return cpu.cb_rl8(cpu.Registers.E)
}
func (cpu *CPU) RL_H() int {
	return cpu.cb_rl8(cpu.Registers.H)
}
func (cpu *CPU) RL_L() int {
	return cpu.cb_rl8(cpu.Registers.L)
}
func (cpu *CPU) RL_HL() int {
	return cpu.cb_rl16(cpu.Registers.HL)
}
func (cpu *CPU) RL_A() int {
	return cpu.cb_rl8(cpu.Registers.A)
}
func (cpu *CPU) RLC_B() int {
	return cpu.cb_rlc8(cpu.Registers.B)
}
func (cpu *CPU) RLC_C() int {
	return cpu.cb_rlc8(cpu.Registers.C)
}
func (cpu *CPU) RLC_D() int {
	return cpu.cb_rlc8(cpu.Registers.D)
}
func (cpu *CPU) RLC_E() int {
	return cpu.cb_rlc8(cpu.Registers.E)
}
func (cpu *CPU) RLC_H() int {
	return cpu.cb_rlc8(cpu.Registers.H)
}
func (cpu *CPU) RLC_L() int {
	return cpu.cb_rlc8(cpu.Registers.L)
}
func (cpu *CPU) RLC_HL() int {
	return cpu.cb_rlc16(cpu.Registers.HL)
}
func (cpu *CPU) RLC_A() int {
	return cpu.cb_rlc8(cpu.Registers.A)
}
func (cpu *CPU) RR_B() int {
	return cpu.cb_rr8(cpu.Registers.B)
}
func (cpu *CPU) RR_C() int {
	return cpu.cb_rr8(cpu.Registers.C)
}
func (cpu *CPU) RR_D() int {
	return cpu.cb_rr8(cpu.Registers.D)
}
func (cpu *CPU) RR_E() int {
	return cpu.cb_rr8(cpu.Registers.E)
}
func (cpu *CPU) RR_H() int {
	return cpu.cb_rr8(cpu.Registers.H)
}
func (cpu *CPU) RR_L() int {
	return cpu.cb_rr8(cpu.Registers.L)
}
func (cpu *CPU) RR_HL() int {
	return cpu.cb_rr16(cpu.Registers.HL)
}
func (cpu *CPU) RR_A() int {
	return cpu.cb_rr8(cpu.Registers.A)
}
func (cpu *CPU) RRC_B() int {
	return cpu.cb_rrc8(cpu.Registers.B)
}
func (cpu *CPU) RRC_C() int {
	return cpu.cb_rrc8(cpu.Registers.C)
}
func (cpu *CPU) RRC_D() int {
	return cpu.cb_rrc8(cpu.Registers.D)
}
func (cpu *CPU) RRC_E() int {
	return cpu.cb_rrc8(cpu.Registers.E)
}
func (cpu *CPU) RRC_H() int {
	return cpu.cb_rrc8(cpu.Registers.H)
}
func (cpu *CPU) RRC_L() int {
	return cpu.cb_rrc8(cpu.Registers.L)
}
func (cpu *CPU) RRC_HL() int {
	return cpu.cb_rrc16(cpu.Registers.HL)
}
func (cpu *CPU) RRC_A() int {
	return cpu.cb_rrc8(cpu.Registers.A)
}
func (cpu *CPU) BIT_0_B() int {
	return cpu.cb_bit(0, cpu.Registers.B)
}

func (cpu *CPU) BIT_0_C() int {
	return cpu.cb_bit(0, cpu.Registers.C)
}

func (cpu *CPU) BIT_0_D() int {
	return cpu.cb_bit(0, cpu.Registers.D)
}

func (cpu *CPU) BIT_0_E() int {
	return cpu.cb_bit(0, cpu.Registers.E)
}

func (cpu *CPU) BIT_0_H() int {
	return cpu.cb_bit(0, cpu.Registers.H)
}

func (cpu *CPU) BIT_0_L() int {
	return cpu.cb_bit(0, cpu.Registers.L)
}

func (cpu *CPU) BIT_0_HL() int {
	return cpu.cb_bit(0, cpu.Registers.HL)
}

func (cpu *CPU) BIT_0_A() int {
	return cpu.cb_bit(0, cpu.Registers.A)
}

func (cpu *CPU) BIT_1_B() int {
	return cpu.cb_bit(1, cpu.Registers.B)
}

func (cpu *CPU) BIT_1_C() int {
	return cpu.cb_bit(1, cpu.Registers.C)
}

func (cpu *CPU) BIT_1_D() int {
	return cpu.cb_bit(1, cpu.Registers.D)
}

func (cpu *CPU) BIT_1_E() int {
	return cpu.cb_bit(1, cpu.Registers.E)
}

func (cpu *CPU) BIT_1_H() int {
	return cpu.cb_bit(1, cpu.Registers.H)
}

func (cpu *CPU) BIT_1_L() int {
	return cpu.cb_bit(1, cpu.Registers.L)
}

func (cpu *CPU) BIT_1_HL() int {
	return cpu.cb_bit(1, cpu.Registers.HL)
}

func (cpu *CPU) BIT_1_A() int {
	return cpu.cb_bit(1, cpu.Registers.A)
}

func (cpu *CPU) BIT_2_B() int {
	return cpu.cb_bit(2, cpu.Registers.B)
}

func (cpu *CPU) BIT_2_C() int {
	return cpu.cb_bit(2, cpu.Registers.C)
}

func (cpu *CPU) BIT_2_D() int {
	return cpu.cb_bit(2, cpu.Registers.D)
}

func (cpu *CPU) BIT_2_E() int {
	return cpu.cb_bit(2, cpu.Registers.E)
}

func (cpu *CPU) BIT_2_H() int {
	return cpu.cb_bit(2, cpu.Registers.H)
}

func (cpu *CPU) BIT_2_L() int {
	return cpu.cb_bit(2, cpu.Registers.L)
}

func (cpu *CPU) BIT_2_HL() int {
	return cpu.cb_bit(2, cpu.Registers.HL)
}

func (cpu *CPU) BIT_2_A() int {
	return cpu.cb_bit(2, cpu.Registers.A)
}

func (cpu *CPU) BIT_3_B() int {
	return cpu.cb_bit(3, cpu.Registers.B)
}

func (cpu *CPU) BIT_3_C() int {
	return cpu.cb_bit(3, cpu.Registers.C)
}

func (cpu *CPU) BIT_3_D() int {
	return cpu.cb_bit(3, cpu.Registers.D)
}

func (cpu *CPU) BIT_3_E() int {
	return cpu.cb_bit(3, cpu.Registers.E)
}

func (cpu *CPU) BIT_3_H() int {
	return cpu.cb_bit(3, cpu.Registers.H)
}

func (cpu *CPU) BIT_3_L() int {
	return cpu.cb_bit(3, cpu.Registers.L)
}

func (cpu *CPU) BIT_3_HL() int {
	return cpu.cb_bit(3, cpu.Registers.HL)
}

func (cpu *CPU) BIT_3_A() int {
	return cpu.cb_bit(3, cpu.Registers.A)
}

func (cpu *CPU) BIT_4_B() int {
	return cpu.cb_bit(4, cpu.Registers.B)
}

func (cpu *CPU) BIT_4_C() int {
	return cpu.cb_bit(4, cpu.Registers.C)
}

func (cpu *CPU) BIT_4_D() int {
	return cpu.cb_bit(4, cpu.Registers.D)
}

func (cpu *CPU) BIT_4_E() int {
	return cpu.cb_bit(4, cpu.Registers.E)
}

func (cpu *CPU) BIT_4_H() int {
	return cpu.cb_bit(4, cpu.Registers.H)
}

func (cpu *CPU) BIT_4_L() int {
	return cpu.cb_bit(4, cpu.Registers.L)
}

func (cpu *CPU) BIT_4_HL() int {
	return cpu.cb_bit(4, cpu.Registers.HL)
}

func (cpu *CPU) BIT_4_A() int {
	return cpu.cb_bit(4, cpu.Registers.A)
}

func (cpu *CPU) BIT_5_B() int {
	return cpu.cb_bit(5, cpu.Registers.B)
}

func (cpu *CPU) BIT_5_C() int {
	return cpu.cb_bit(5, cpu.Registers.C)
}

func (cpu *CPU) BIT_5_D() int {
	return cpu.cb_bit(5, cpu.Registers.D)
}

func (cpu *CPU) BIT_5_E() int {
	return cpu.cb_bit(5, cpu.Registers.E)
}

func (cpu *CPU) BIT_5_H() int {
	return cpu.cb_bit(5, cpu.Registers.H)
}

func (cpu *CPU) BIT_5_L() int {
	return cpu.cb_bit(5, cpu.Registers.L)
}

func (cpu *CPU) BIT_5_HL() int {
	return cpu.cb_bit(5, cpu.Registers.HL)
}

func (cpu *CPU) BIT_5_A() int {
	return cpu.cb_bit(5, cpu.Registers.A)
}

func (cpu *CPU) BIT_6_B() int {
	return cpu.cb_bit(6, cpu.Registers.B)
}

func (cpu *CPU) BIT_6_C() int {
	return cpu.cb_bit(6, cpu.Registers.C)
}

func (cpu *CPU) BIT_6_D() int {
	return cpu.cb_bit(6, cpu.Registers.D)
}

func (cpu *CPU) BIT_6_E() int {
	return cpu.cb_bit(6, cpu.Registers.E)
}

func (cpu *CPU) BIT_6_H() int {
	return cpu.cb_bit(6, cpu.Registers.H)
}

func (cpu *CPU) BIT_6_L() int {
	return cpu.cb_bit(6, cpu.Registers.L)
}

func (cpu *CPU) BIT_6_HL() int {
	return cpu.cb_bit(6, cpu.Registers.HL)
}

func (cpu *CPU) BIT_6_A() int {
	return cpu.cb_bit(6, cpu.Registers.A)
}

func (cpu *CPU) BIT_7_B() int {
	return cpu.cb_bit(7, cpu.Registers.B)
}

func (cpu *CPU) BIT_7_C() int {
	return cpu.cb_bit(7, cpu.Registers.C)
}

func (cpu *CPU) BIT_7_D() int {
	return cpu.cb_bit(7, cpu.Registers.D)
}

func (cpu *CPU) BIT_7_E() int {
	return cpu.cb_bit(7, cpu.Registers.E)
}

func (cpu *CPU) BIT_7_H() int {
	return cpu.cb_bit(7, cpu.Registers.H)
}

func (cpu *CPU) BIT_7_L() int {
	return cpu.cb_bit(7, cpu.Registers.L)
}

func (cpu *CPU) BIT_7_HL() int {
	return cpu.cb_bit(7, cpu.Registers.HL)
}

func (cpu *CPU) BIT_7_A() int {
	return cpu.cb_bit(7, cpu.Registers.A)
}

func (cpu *CPU) SWAP_B() int {
	return cpu.cb_swap(cpu.Registers.B)
}

func (cpu *CPU) SWAP_C() int {
	return cpu.cb_swap(cpu.Registers.C)
}

func (cpu *CPU) SWAP_D() int {
	return cpu.cb_swap(cpu.Registers.D)
}

func (cpu *CPU) SWAP_E() int {
	return cpu.cb_swap(cpu.Registers.E)
}

func (cpu *CPU) SWAP_H() int {
	return cpu.cb_swap(cpu.Registers.H)
}

func (cpu *CPU) SWAP_L() int {
	return cpu.cb_swap(cpu.Registers.L)
}

func (cpu *CPU) SWAP_HL() int {
	return cpu.cb_swap(cpu.Registers.HL)
}

func (cpu *CPU) SWAP_A() int {
	return cpu.cb_swap(cpu.Registers.A)
}

func (cpu *CPU) SRL_B() int {
	return cpu.cb_srl(cpu.Registers.B)
}

func (cpu *CPU) SRL_C() int {
	return cpu.cb_srl(cpu.Registers.C)
}

func (cpu *CPU) SRL_D() int {
	return cpu.cb_srl(cpu.Registers.D)
}

func (cpu *CPU) SRL_E() int {
	return cpu.cb_srl(cpu.Registers.E)
}

func (cpu *CPU) SRL_H() int {
	return cpu.cb_srl(cpu.Registers.H)
}

func (cpu *CPU) SRL_L() int {
	return cpu.cb_srl(cpu.Registers.L)
}

func (cpu *CPU) SRL_HL() int {
	return cpu.cb_srl(cpu.Registers.HL)
}

func (cpu *CPU) SRL_A() int {
	return cpu.cb_srl(cpu.Registers.A)
}
