package main

type cbFun func(byte) byte

func (cpu *CPU) Oxcb() int {
	switch opcode := cpu.FetchAndIncrement(); opcode {
	case 0x7c:
		return cpu.BIT_7H()
	case 0x11:
		return cpu.RL_C()
	}

	return 8

}

func (cpu *CPU) BIT_7H() int {
	// copy the contents of of bit 7 in register H to the z flag of the F register.
	bit := cpu.Registers.H.GetBit(7) ^ 1 // take complemeent of the bit in position 7.
	cpu.Registers.F.SetBit(bit, Z_FLAG)
	cpu.Registers.F.SetBit(1, HALF_CARRY_FLAG)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	return 4
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
