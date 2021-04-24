package main

type cbFun func(byte) byte

func (cpu *CPU) Oxcb() int {
	switch opcode := cpu.FetchAndIncrement(); opcode {
	case 0x7c:
		return cpu.BIT_7_H()
	case 0x11:
		return cpu.RL_C()
	}

	return 8
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

func (cpu *CPU) RL_HL() int {
	return cpu.cb_rl16(cpu.Registers.HL)
}

func (cpu *CPU) RL_A() int {
	return cpu.cb_rl8(cpu.Registers.A)
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

func (cpu *CPU) RR_HL() int {
	return cpu.cb_rr16(cpu.Registers.HL)
}

func (cpu *CPU) RR_A() int {
	return cpu.cb_rr8(cpu.Registers.A)
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

func (cpu *CPU) RLC_HL() int {
	return cpu.cb_rlc16(cpu.Registers.HL)
}

func (cpu *CPU) RLC_A() int {
	return cpu.cb_rlc8(cpu.Registers.A)
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

func (cpu *CPU) RRC_HL() int {
	return cpu.cb_rrc16(cpu.Registers.HL)
}

func (cpu *CPU) RRC_A() int {
	return cpu.cb_rrc8(cpu.Registers.A)
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

func (cpu *CPU) SRA_HL() int {
	return cpu.cb_sra16(cpu.Registers.HL)
}

func (cpu *CPU) SRA_A() int {
	return cpu.cb_sra8(cpu.Registers.A)
}

func (cpu *CPU) swap(val byte) byte {
	bits := make([]byte, 8)
	for i := 0; i < 8; i++  {
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

func (cpu *CPU) SWAP_HL() int {
	return cpu.cb_swap(cpu.Registers.HL)
}

func (cpu *CPU) SWAP_A() int {
	return cpu.cb_swap(cpu.Registers.A)
}

func (cpu *CPU) cb_srl(register interface{}) int {
	cpu.Registers.F.SetBit(0, CARRY_FLAG) // ensuring bit 7 is set to 0.
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

func (cpu *CPU) SRL_HL() int {
	return cpu.cb_srl(cpu.Registers.HL)
}

func (cpu *CPU) SRL_A() int {
	return cpu.cb_srl(cpu.Registers.A)
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

func (cpu *CPU) BIT_7_HL() int {
	return cpu.cb_bit(7, cpu.Registers.HL)
}

func (cpu *CPU) BIT_7_A() int {
	return cpu.cb_bit(7, cpu.Registers.A)
}

func (cpu *CPU) cb_res_set(index int, bit byte,  register interface{}) int {
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

func (cpu *CPU) SET_7_HL() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.HL)
}

func (cpu *CPU) SET_7_A() int {
	return cpu.cb_res_set(7, 1, cpu.Registers.A)
}
