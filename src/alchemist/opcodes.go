package main

import "fmt"

func (cpu *CPU) LD_SP_D16(){
	// 0x31
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.SP = MergeBytes(high, low)
}
func (cpu *CPU) XOR_A() {
	// xor A
	cpu.Registers.A.Set(cpu.Registers.A.Value ^ cpu.Registers.A.Value)
}

func (cpu *CPU) LD_HL_A() {
	// LD_HL_A
	a := cpu.Registers.A.Get()
	cpu.Memory[cpu.Registers.HL.Get()] = a
	cpu.Registers.HL.Decrement()
}


func (cpu *CPU) LD_HL_D16() {
	// 0x21
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.Registers.H.Set(high)
	cpu.Registers.L.Set(low)
}





func (cpu *CPU) JR_NZ_S8(){
	zFlag := cpu.Registers.F.GetBit(Z_FLAG)
	pc := &cpu.PC
	*pc += 1
	if zFlag == 0 {
		steps := GetTwosComplement(cpu.Memory[*pc])
		fmt.Println(*pc)
		fmt.Println(int(*pc) + steps)
	}
}