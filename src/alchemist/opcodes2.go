package main

import "fmt"

func (cpu *CPU) Oxcb() {
	switch opcode := cpu.FetchAndIncrement() ; opcode {
	case 0x7c:
		cpu.BIT_7H()
	}

}

func (cpu *CPU) BIT_7H()  {
	// copy the contents of of bit 7 in register H to the z flag of the F register.
	fmt.Println(fmt.Sprintf("H Binary : %b", cpu.Registers.H.Get()))
	fmt.Println(fmt.Sprintf("HL Binary %b", cpu.Registers.HL.Get()))
	bit := cpu.Registers.H.GetBit(7) ^ 1 // take complemeent of the bit in position 7.
	cpu.Registers.F.SetBit(bit, Z_FLAG)
}


