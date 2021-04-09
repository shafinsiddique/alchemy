package main

func (cpu *CPU) Oxcb() {
	switch opcode := cpu.FetchAndIncrement() ; opcode {
	case 0x7c:
		cpu.BIT_7H()
	case 0x11:
		cpu.RL_C()
	}

}

func (cpu *CPU) BIT_7H()  {
	// copy the contents of of bit 7 in register H to the z flag of the F register.
	bit := cpu.Registers.H.GetBit(7) ^ 1 // take complemeent of the bit in position 7.
	cpu.Registers.F.SetBit(bit, Z_FLAG)
	cpu.Registers.F.SetBit(1, HALF_CARRY_FLAG)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
}

func (cpu *CPU) RL_C(){
	cpu.RL(cpu.Registers.C)
	cpu.CheckAndSetZeroFlag(cpu.Registers.C.Get())
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
}

func (cpu *CPU) SUB_B() {
	cpu.Registers.A.Set(cpu.Registers.A.Get()-cpu.Registers.B.Get())
}


