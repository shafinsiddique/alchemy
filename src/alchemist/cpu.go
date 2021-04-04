package main

import "fmt"

type CPU struct {
	Registers *Registers
	Memory []byte
	PC uint16
	SP uint16
}

func NewCPU() *CPU {
	return &CPU{Registers: InitializeRegisters(), Memory: make([]byte,0x10000)}
}

func (cpu *CPU) LoadBootRom(bootrom []byte) {
	for i := 0; i < len(bootrom) ; i++ {
		cpu.Memory[i] = bootrom[i]
	}
}

func (cpu *CPU) ldSpD16(){
	// 0x31
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.SP = MergeBytes(high, low)
}
func (cpu *CPU) xorA() {
	// xor A
	cpu.Registers.A.Set(cpu.Registers.A.Value ^ cpu.Registers.A.Value)
}

func (cpu *CPU) ldHlA() {
	// ldHlA
	a := cpu.Registers.A.Get()
	cpu.Memory[cpu.Registers.HL.Get()] = a
	cpu.Registers.HL.Decrement()
}


func (cpu *CPU) ldHlD16() {
	// 0x21
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.Registers.H.Set(high)
	cpu.Registers.L.Set(low)
}

func (cpu *CPU) bit7H()  {
	// copy the contents of of bit 7 in register H to the z flag of the F register.
	bit := cpu.Registers.H.GetBit(7) ^ 1 // take complemeent of the bit in position 7.
	cpu.Registers.F.SetBit(bit, Z_FLAG)
}

func (cpu *CPU) Oxcb() {
	pc := &cpu.PC
	*pc += 1
	switch opcode := cpu.Memory[*pc] ; opcode {
	case 0x7c:
		cpu.bit7H()
	}

}

func (cpu *CPU) jrNzS8(){
	zFlag := cpu.Registers.F.GetBit(Z_FLAG)
	pc := &cpu.PC
	*pc += 1
	if zFlag == 0 {
		steps := GetTwosComplement(cpu.Memory[*pc])
		fmt.Println(*pc)
		fmt.Println(int(*pc) + steps)
	}
}

func (cpu *CPU) FetchDecodeExecute() {
	pc := &cpu.PC
	switch opcode := cpu.Memory[*pc]; opcode {
	case 0x31:
		cpu.ldSpD16()
	case 0xaf:
		cpu.xorA()
	case 0x21:
		cpu.ldHlD16()
	case 0x32:
		cpu.ldHlA()
	case 0xcb:
		cpu.Oxcb()
	case 0x20:
		fmt.Println("here")
		hex := fmt.Sprintf("%x", opcode)
		fmt.Println("0x" + hex)
		cpu.jrNzS8() // s8 stands for signed 8 bit.
		fmt.Println("finished")

	default:
		hex := fmt.Sprintf("%x", opcode)
		fmt.Println("0x" + hex)
	}
	*pc += 1 // always increment one, even if other instructions increment, we need to increment from that position to
	// go to the next one. basically, this allows us to not worry abouyt incrementing one at the end of every single
	// function.

}

func (cpu *CPU) RunBootSequence(){
	for cpu.PC < 256 {
		cpu.FetchDecodeExecute()
	}
}

