package main

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

func (cpu *CPU) LD_HL_A_DEC() {
	// LD_HL_A_DEC
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
		*pc += int16(steps)
	}
}

func (cpu *CPU) LD_C_D8() {
	pc := &cpu.PC
	*pc += 1
	nextByte := cpu.Memory[*pc]
	cpu.Registers.C.Set(nextByte)
}

func (cpu *CPU) LD_A_D8(){
	pc := &cpu.PC
	*pc += 1
	nextByte := cpu.Memory[*pc]
	cpu.Registers.A.Set(nextByte)
}

func (cpu *CPU) LD_LOC_C_A() {
	// store the contents of register A in the internal ram, ad the range 0xff00-0xffff specified by register c.
	// disassembly in boot rom : LD (0xFF00 + C), A
	addr := 0xff + cpu.Registers.C.Get()
	cpu.Memory[addr] = cpu.Registers.A.Get()
}

func (cpu *CPU) INC_C() {
	cpu.Registers.C.Set(cpu.Registers.C.Get() + 1)
}

func (cpu *CPU) LD_HL_A() {
	// store the contents of register a in the memory location specified by HL
	cpu.Memory[cpu.Registers.HL.Get()] = cpu.Registers.A.Get()
}

func (cpu *CPU) LD_A8_A() {
	// store the contents of register A in the range 0xFF00-0xFFf specified by immediarte
	// operand a8.

	pc := &cpu.PC
	*pc += 1
	addr := 0xff + cpu.Memory[*pc]
	cpu.Memory[addr] = cpu.Registers.A.Get()
}

func (cpu *CPU) LD_DE_D16() {
	// load the 2 bytes of immediate data into register pair DE.
	pc := &cpu.PC
	*pc += 1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]
	cpu.Registers.D.Set(high)
	cpu.Registers.E.Set(low)
}

func (cpu *CPU) LD_A_LOC_DE() {
	// store the 8 bit contents in the memory location of the value of DE into register A.
	cpu.Registers.A.Set(cpu.Memory[cpu.Registers.DE.Get()])
}

func (cpu *CPU) CALL_A16(){
	// push the program counter PC value corresponding to the address following the CALL instruction.
	// TO the 2 bytes following the byte specified by the current statck pointer SP.
	// Then load the 16 bit immediate operand a16 into PC.
	pc := &cpu.PC // part 1 store PC following this instruction in stack pointer.
	sp := &cpu.SP
	*sp -= 1
	bytes := SplitInt16ToBytes(uint16(*pc)+3) // + 3 because length of instruction in
	cpu.Memory[*sp] = bytes[0] // high byte placed on top of low byte.
	*sp -= 1
	cpu.Memory[*sp] = bytes[1] // low byte placed bottom , i guess the name makes sense?

	// part ii load 16 bit immediate operand.
	*pc +=1
	low := cpu.Memory[*pc]
	*pc += 1
	high := cpu.Memory[*pc]

	*pc = int16(MergeBytes(high, low)) - 1 // minus 1 because in the fetchExecuteDecode method, the pc will
	// be incremented once this function returns.
}

func (cpu *CPU) LD_C_A() {
	cpu.Registers.A.Set(cpu.Registers.C.Get())
}

func (cpu *CPU) LD_B_D8(){
	// ld 8 bit immediate into register b.
	pc := &cpu.PC
	*pc += 1
	cpu.Registers.B.Set(cpu.Memory[*pc])
}

func (cpu *CPU) PUSH_BC() {
	// push contents of bc into stack.
	cpu.PushToStack(cpu.Registers.BC.GetHigh())
	cpu.PushToStack(cpu.Registers.BC.GetLow())
}

func (cpu *CPU) RLA(){
	// rotate the contents of reigster a to the left, trhough the carry flag.
	// i.e bit 0 -> bit 1 -> bit 2
	for i:= 0; i < 8; i++ {
		bit := cpu.Registers.A.GetBit(i)
		carry := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
		cpu.Registers.A.SetBit(carry, i)

	}
}

func (cpu *CPU) POP_BC() {
	// pop the contents from the memory stack into BC.
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	cpu.Registers.B.Set(high)
	cpu.Registers.C.Set(low)
}

func (cpu *CPU) DEC_B() {
	cpu.Registers.B.Decrement()
}

func (cpu *CPU) LD_LOC_HL_A_INC() {
	// store the element in memory loc HL into register A.
	// also increment HL.

	cpu.Registers.A.Set(cpu.Memory[cpu.Registers.HL.Get()])
	cpu.Registers.HL.Increment()
}

func (cpu *CPU) INC_HL() {
	cpu.Registers.HL.Increment()
}

func (cpu *CPU) RET(){
	// pop from the stack the PC value pushed when subroutine was called.
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	cpu.PC = int16(MergeBytes(high, low) - 1) // -1 for now because the pc will be incremement. TODO: fix
}




