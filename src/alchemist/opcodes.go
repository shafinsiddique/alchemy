package main

import "fmt"

func (cpu *CPU) LD_SP_D16(){
	// 0x31
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
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
	fmt.Println("Decrementing HL : ", cpu.Registers.HL.Get())
}


func (cpu *CPU) LD_HL_D16() {
	// 0x21
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Registers.H.Set(high)
	cpu.Registers.L.Set(low)
	fmt.Println("Setting HL : ", cpu.Registers.HL.Get())

}

func (cpu *CPU) JR_COMMON_S8(flag byte) { // not actual instreuction, code reuse.
	zFlag := cpu.Registers.F.GetBit(Z_FLAG)
	nextByte := cpu.FetchAndIncrement()
	if zFlag == flag {
		steps, isNegative := GetTwosComplement(nextByte)
		if isNegative {
			cpu.PC -= uint16(steps)
		} else {
			cpu.PC += uint16(steps)
		}
	}
}

func (cpu *CPU) JR_NZ_S8(){
	cpu.JR_COMMON_S8(0)
}

func (cpu *CPU) LD_C_D8() {
	nextByte := cpu.FetchAndIncrement()
	cpu.Registers.C.Set(nextByte)
}

func (cpu *CPU) LD_A_D8(){
	nextByte := cpu.FetchAndIncrement()
	cpu.Registers.A.Set(nextByte)
}

func (cpu *CPU) LD_LOC_C_A() {
	// store the contents of register A in the internal ram, ad the range 0xff00-0xffff specified by register c.
	// disassembly in boot rom : LD (0xFF00 + C), A
	addr := 0xff00 + uint16(cpu.Registers.C.Get())
	cpu.Memory[addr] = cpu.Registers.A.Get()
}

func (cpu *CPU) INC_C() {
	cpu.Registers.C.Set(cpu.Registers.C.Get() + 1)
}

func (cpu *CPU) LD_LOC_HL_A() {
	// store the contents of register a in the memory location specified by HL
	cpu.Memory[cpu.Registers.HL.Get()] = cpu.Registers.A.Get()
}

func (cpu *CPU) LD_LOC_A8_A() {
	// store the contents of register A in the range 0xFF00-0xFFf specified by immediarte
	// operand a8.
	addr := 0xff00 + uint16(cpu.FetchAndIncrement())
	cpu.Memory[addr] = cpu.Registers.A.Get()
}

func (cpu *CPU) LD_DE_D16() {
	// load the 2 bytes of immediate data into register pair DE.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
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
	sp := &cpu.SP
	*sp -= 1
	bytes := SplitInt16ToBytes(uint16(cpu.PC + 2)) // + 2 because current PC = Position of Call + 1
	cpu.Memory[*sp] = bytes[0]                     // current byte and next byte is included so + 2 goes to next
	*sp -= 1										// instruction
	cpu.Memory[*sp] = bytes[1] // low byte placed bottom , i guess the name makes sense?

	// part ii load 16 bit immediate operand.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()

	cpu.PC = MergeBytes(high, low)
	// be incremented once this function returns.
}

func (cpu *CPU) LD_C_A() {
	cpu.Registers.A.Set(cpu.Registers.C.Get())
}

func (cpu *CPU) LD_B_D8(){
	// ld 8 bit immediate into register b.
	cpu.Registers.B.Set(cpu.FetchAndIncrement())
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
	cpu.PC = MergeBytes(high, low)
}

func (cpu *CPU) INC_DE() {
	cpu.Registers.DE.Increment()
}

func (cpu *CPU) LD_A_E() {
	// load the contents of register E into register A.
	cpu.Registers.A.Set(cpu.Registers.E.Get())
}

func (cpu *CPU) CP_D8() {
	// compare the contents of reigster A with immediate 8 bit operand d8. set z flag if they are
	// equal.
	i := cpu.FetchAndIncrement()
	if cpu.Registers.A.Get() - i == 0 {
		cpu.Registers.F.SetBit(1, Z_FLAG)
	}
}

func (cpu *CPU) LD_LOC_A16_A(){
	// store the contents of register A in the internal ram specified by the 16 bit immeidate operand a16.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Memory[MergeBytes(high, low)] = cpu.Registers.A.Get()
}

func (cpu *CPU) DEC_A() {
	cpu.Registers.A.Decrement()
}

func (cpu *CPU) JR_Z_S8() {
	cpu.JR_COMMON_S8(1)
}

func (cpu *CPU) LD_H_A() {
	cpu.Registers.H.Set(cpu.Registers.A.Get())
}

func (cpu *CPU) LD_D_A() {
	cpu.Registers.D.Set(cpu.Registers.A.Get())
}

func (cpu *CPU) INC_B() {
	cpu.Registers.B.Increment()
}

func (cpu *CPU) LD_E_D8() {
	cpu.Registers.E.Set(cpu.FetchAndIncrement())
}

func (cpu *CPU) LD_A_LOC_A8() {
	addr := 0xff00 + uint16(cpu.FetchAndIncrement())
	cpu.Registers.A.Set(cpu.Memory[addr])
}

func (cpu *CPU) DEC_C(){
	cpu.Registers.C.Decrement()
}

func (cpu *CPU) DEC_E() {
	cpu.Registers.E.Decrement()
}

func (cpu *CPU) DEC_D(){
	cpu.Registers.D.Decrement()
}

func (cpu *CPU) LD_D_D8(){
	cpu.Registers.D.Set(cpu.FetchAndIncrement())
}

func (cpu *CPU) INC_H(){
	cpu.Registers.H.Increment()
}


