package main

import "fmt"

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

	if bit == 1 {
		fmt.Println(cpu.Registers.HL.Get())
		fmt.Println(cpu.Registers.F.GetBit(Z_FLAG))
		fmt.Println("hello world.")
	}
}

func (cpu *CPU) RL_C(){
	cpu.RL(cpu.Registers.C)
}

func (cpu *CPU) SUB_B() {
	cpu.Registers.A.Set(cpu.Registers.A.Get()-cpu.Registers.B.Get())
}


