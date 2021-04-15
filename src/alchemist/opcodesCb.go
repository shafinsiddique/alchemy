package main

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

func (cpu *CPU) RL_C() int {
	cpu.CheckAndSetZeroFlag(cpu.Registers.C.Get())
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
	return cpu.RL(cpu.Registers.C)

}

func (cpu *CPU) SUB_B() int {
	a := cpu.Registers.A.Get()
	b := cpu.Registers.B.Get()

	if !cpu.CheckAndSetOverflowFlag(a, b, true)  {
		diff := a-b
		cpu.Registers.A.Set(diff)
		cpu.CheckAndSetZeroFlag(diff)

	} else { // since it cant be
		cpu.ClearZeroFlag()
	}

	cpu.CheckAndSetHCFlag(a, b, true)
	cpu.SetNegativeFlag()
	return 4
}
