package main

func (cpu *CPU) LD_SP_D16() int {
	// 0x31
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.SP = MergeBytes(high, low)
	return 12
}

func (cpu *CPU) XOR_A() int {
	// xor A
	cpu.Registers.A.Set(cpu.Registers.A.Value ^ cpu.Registers.A.Value)
	cpu.CheckAndSetZeroFlag(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) LD_LOC_HL_A_DEC() int {
	// LD_LOC_HL_A_DEC
	a := cpu.Registers.A.Get()
	cpu.MMU.Write(cpu.Registers.HL.Get(), a)
	cpu.Registers.HL.Decrement()
	return 8
}

func (cpu *CPU) LD_HL_D16() int {
	// 0x21
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Registers.H.Set(high)
	cpu.Registers.L.Set(low)
	return 12
}

func (cpu *CPU) JR_COMMON_S8(flag byte) int { // not actual instreuction, code reuse.
	zFlag := cpu.Registers.F.GetBit(Z_FLAG)
	nextByte := cpu.FetchAndIncrement()
	cycles := 8 // Cycles without branch,
	if zFlag == flag {
		steps, isNegative := GetTwosComplement(nextByte)
		if isNegative {
			cpu.PC -= uint16(steps)
		} else {
			cpu.PC += uint16(steps)
		}
		cycles = 12
	}

	return cycles
}

func (cpu *CPU) RL(register *EightBitRegister) int {
	for i := 0; i <= 7; i++ {
		bit := register.GetBit(i)
		carry := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
		register.SetBit(carry, i)
	}
	return 4
}

func (cpu *CPU) JR_NZ_S8() int {
	return cpu.JR_COMMON_S8(0)
}

func (cpu *CPU) LD_C_D8() int {
	nextByte := cpu.FetchAndIncrement()
	cpu.Registers.C.Set(nextByte)
	return 8
}

func (cpu *CPU) LD_A_D8() int {
	nextByte := cpu.FetchAndIncrement()
	cpu.Registers.A.Set(nextByte)
	return 8
}

func (cpu *CPU) LD_LOC_C_A() int {
	// store the contents of register A in the internal ram, ad the range 0xff00-0xffff specified by register c.
	// disassembly in boot rom : LD (0xFF00 + C), A
	addr := 0xff00 + uint16(cpu.Registers.C.Get())
	cpu.MMU.Write(addr, cpu.Registers.A.Get())
	return 8
}

func (cpu *CPU) INC_C() int {
	cpu.IncrementRegister8Bit(cpu.Registers.C)
	return 4
}

func (cpu *CPU) LD_LOC_HL_A() int {
	// store the contents of register a in the memory location specified by HL
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.A.Get())
	return 8
}

func (cpu *CPU) LD_LOC_A8_A() int {
	// store the contents of register A in the range 0xFF00-0xFFf specified by immediarte
	// operand a8.
	addr := 0xff00 + uint16(cpu.FetchAndIncrement())
	cpu.MMU.Write(addr, cpu.Registers.A.Get())
	return 12
}

func (cpu *CPU) LD_DE_D16() int {
	// load the 2 bytes of immediate data into register pair DE.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Registers.D.Set(high)
	cpu.Registers.E.Set(low)
	return 12
}

func (cpu *CPU) LD_A_LOC_DE() int {
	// store the 8 bit contents in the memory location of the value of DE into register A.
	cpu.Registers.A.Set(cpu.MMU.Read(cpu.Registers.DE.Get()))
	return 8
}

func (cpu *CPU) CALL_A16() int {
	// push the program counter PC value corresponding to the address following the CALL instruction.
	// TO the 2 bytes following the byte specified by the current statck pointer SP.
	// Then load the 16 bit immediate operand a16 into PC.
	sp := &cpu.SP
	*sp -= 1
	bytes := SplitInt16ToBytes(uint16(cpu.PC + 2)) // + 2 because current PC = Position of Call + 1
	cpu.MMU.Write(*sp, bytes[0])                   // high byte placed at the top.
	*sp -= 1
	cpu.MMU.Write(*sp, bytes[1]) // low byte placed bottom , i guess the name makes sense?

	// part ii load 16 bit immediate operand.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()

	cpu.PC = MergeBytes(high, low)
	// be incremented once this function returns.
	return 24
}

func (cpu *CPU) LD_C_A() int {
	cpu.Registers.C.Set(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) LD_B_D8() int {
	// ld 8 bit immediate into register b.
	cpu.Registers.B.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) PUSH_BC() int {
	// push contents of bc into stack.
	cpu.PushToStack(cpu.Registers.BC.GetHigh())
	cpu.PushToStack(cpu.Registers.BC.GetLow())
	return 16
}

func (cpu *CPU) RL_A() int {
	// rotate the contents of reigster a to the left, trhough the carry flag.
	// i.e bit 0 -> bit 1 -> bit 2
	cpu.Registers.F.SetBit(0, Z_FLAG)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
	return cpu.RL(cpu.Registers.A)

}

func (cpu *CPU) POP_BC() int {
	// pop the contents from the memory stack into BC.
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	cpu.Registers.B.Set(high)
	cpu.Registers.C.Set(low)
	return 12
}

func (cpu *CPU) DEC_B() int {
	return cpu.DecrementRegister8Bit(cpu.Registers.B)
}

func (cpu *CPU) LD_LOC_HL_A_INC() int {
	// store the element in memory loc HL into register A.
	// also increment HL.
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.A.Get())
	cpu.Registers.HL.Increment()
	return 12
}

func (cpu *CPU) INC_HL() int {
	cpu.Registers.HL.Increment()
	return SIXTEEN_BIT_INC_CYCLE
}

func (cpu *CPU) RET() int {
	// pop from the stack the PC value pushed when subroutine was called.
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	cpu.PC = MergeBytes(high, low)
	return 16
}

func (cpu *CPU) INC_DE() int {
	cpu.Registers.DE.Increment()
	return SIXTEEN_BIT_INC_CYCLE
}

func (cpu *CPU) LD_A_E() int {
	// load the contents of register E into register A.
	cpu.Registers.A.Set(cpu.Registers.E.Get())
	return LD_CYCLE
}

func (cpu *CPU) CP_D8() int {
	// compare the contents of reigster A with immediate 8 bit operand d8. set z flag if they are
	// equal.
	return cpu.CP(cpu.FetchAndIncrement())
}

func (cpu *CPU) LD_LOC_A16_A() int {
	// store the contents of register A in the internal ram specified by the 16 bit immeidate operand a16.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.MMU.Write(MergeBytes(high, low), cpu.Registers.A.Get())
	return LD_SIXTEEN_BIT_CYCLE
}

func (cpu *CPU) DEC_A() int {
	return cpu.DecrementRegister8Bit(cpu.Registers.A)
}

func (cpu *CPU) JR_Z_S8() int {
	return cpu.JR_COMMON_S8(1)
}

func (cpu *CPU) LD_H_A() int {
	cpu.Registers.H.Set(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) INC_B() int {
	return cpu.IncrementRegister8Bit(cpu.Registers.B)
}

func (cpu *CPU) LD_E_D8() int {
	cpu.Registers.E.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) LD_A_LOC_A8() int {
	addr := 0xff00 + uint16(cpu.FetchAndIncrement())
	cpu.Registers.A.Set(cpu.MMU.Read(addr))
	return 12
}

func (cpu *CPU) DEC_C() int {
	return cpu.DecrementRegister8Bit(cpu.Registers.C)
}

func (cpu *CPU) DEC_E() int {
	return cpu.DecrementRegister8Bit(cpu.Registers.E)
}

func (cpu *CPU) DEC_D() int {
	return cpu.DecrementRegister8Bit(cpu.Registers.D)
}

func (cpu *CPU) LD_D_D8() int {
	cpu.Registers.D.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) INC_H() int {
	return cpu.IncrementRegister8Bit(cpu.Registers.H)
}

func (cpu *CPU) LD_A_H() int {
	cpu.Registers.A.Set(cpu.Registers.H.Get())
	return 4
}

func (cpu *CPU) JR_S8() int {
	// jump s8 steps from current.

	steps, isN := GetTwosComplement(cpu.FetchAndIncrement())
	steps16 := uint16(steps)
	if isN {
		cpu.PC -= steps16
	} else {
		cpu.PC += steps16
	}
	return 12
}

func (cpu *CPU) LD_L_D8() int {
	cpu.Registers.L.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) CP(val byte) int {
	// compare the contents of register A with Val by calculating A-(HL)
	a := cpu.Registers.A.Get()
	if !cpu.CheckAndSetOverflowFlag(a, val, true) {
		cpu.CheckAndSetZeroFlag(a - val)
	} else { // since the a is definitely not zero/
		cpu.ClearZeroFlag()
	}
	cpu.SetNegativeFlag()
	cpu.CheckAndSetHCFlag(a, val, true)
	return 8
}

func (cpu *CPU) CP_LOC_HL() int {
	return cpu.CP(cpu.MMU.Read(cpu.Registers.HL.Get()))
}

func (cpu *CPU) LD_A_L() int {
	cpu.Registers.A.Set(cpu.Registers.L.Get())
	return 4
}

func (cpu *CPU) LD_A_B() int {
	cpu.Registers.A.Set(cpu.Registers.B.Get())
	return 4
}

func (cpu *CPU) ADD_AND_STORE(register *EightBitRegister, val2 byte) {
	val1 := register.Get()
	sum := val1 + val2
	register.Set(sum)
	cpu.CheckAndSetZeroFlag(sum)
	cpu.ClearNegativeFlag()
	cpu.CheckAndSetHCFlag(val1, val2, false)
	cpu.CheckAndSetOverflowFlag(val1, val2, false)
}

func (cpu *CPU) ADD_A_LOC_HL() int {
	cpu.ADD_AND_STORE(cpu.Registers.A, cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) LD_BC_D16() int {
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Registers.BC.Set(high, low)
	return 12
}

func (cpu *CPU) LD_LOC_BC_A() int {
	cpu.MMU.Write(cpu.Registers.BC.Get(), cpu.Registers.A.Get())
	return 8
}

func (cpu *CPU) LD_LOC_DE_A() int {
	cpu.MMU.Write(cpu.Registers.DE.Get(), cpu.Registers.A.Get())
	return 8
}

func (cpu *CPU) LD_B_B() int {
	cpu.Registers.B.Set(cpu.Registers.B.Get())
	return 4
}

func (cpu *CPU) LD_D_B() int {
	cpu.Registers.D.Set(cpu.Registers.B.Get())
	return 4
}

func (cpu *CPU) LD_H_B() int {
	cpu.Registers.H.Set(cpu.Registers.B.Get())
	return 4
}

func (cpu *CPU) LD_LOC_HL_B() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.B.Get())
	return 8
}

func (cpu *CPU) LD_LOC_HL_C() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.C.Get())
	return 8
}

func (cpu *CPU) LD_LOC_HL_D() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.D.Get())
	return 8
}

func (cpu *CPU) LD_LOC_HL_E() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.E.Get())
	return 8
}

func (cpu *CPU) LD_LOC_HL_H() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.H.Get())
	return 8
}

func (cpu *CPU) LD_LOC_HL_L() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.Registers.L.Get())
	return 8
}

func (cpu *CPU) LD_A_C() int {
	cpu.Registers.A.Set(cpu.Registers.C.Get())
	return 4
}

func (cpu *CPU) LD_A_D() int {
	cpu.Registers.A.Set(cpu.Registers.D.Get())
	return 4
}

func (cpu *CPU) LD_A_LOC_HL() int {
	cpu.Registers.A.Set(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) LD_A_A() int {
	cpu.Registers.A.Set(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) LD_L_A() int {
	cpu.Registers.L.Set(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) LD_L_LOC_HL() int {
	cpu.Registers.L.Set(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) LD_L_L() int {
	cpu.Registers.L.Set(cpu.Registers.L.Get())
	return 4
}

func (cpu *CPU) LD_L_H() int {
	cpu.Registers.L.Set(cpu.Registers.H.Get())
	return 4
}

func (cpu *CPU) LD_L_E() int {
	cpu.Registers.L.Set(cpu.Registers.E.Get())
	return 4
}

func (CPU *CPU) ld_dst_src(dst *EightBitRegister, src *EightBitRegister) int {
	dst.Set(src.Get())
	return 4
}

func (cpu *CPU) ld_dst_loc_src(dst *EightBitRegister, src *SixteenBitRegister) int {
	dst.Set(cpu.MMU.Read(src.Get()))
	return 8
}

func (cpu *CPU) LD_L_D() int {
	return cpu.ld_dst_src(cpu.Registers.L, cpu.Registers.D)
}

func (cpu *CPU) LD_L_C() int {
	return cpu.ld_dst_src(cpu.Registers.L, cpu.Registers.C)
}

func (cpu *CPU) LD_L_B() int {
	return cpu.ld_dst_src(cpu.Registers.L, cpu.Registers.B)
}

func (cpu *CPU) LD_H_LOC_HL() int {
	return cpu.ld_dst_loc_src(cpu.Registers.H, cpu.Registers.HL)
}

func (cpu *CPU) LD_H_L() int {
	return cpu.ld_dst_src(cpu.Registers.H, cpu.Registers.L)
}

func (cpu *CPU) LD_H_H() int {
	return cpu.ld_dst_src(cpu.Registers.H, cpu.Registers.H)
}

func (cpu *CPU) LD_H_E() int {
	return cpu.ld_dst_src(cpu.Registers.H, cpu.Registers.E)
}

func (cpu *CPU) LD_H_D() int {
	return cpu.ld_dst_src(cpu.Registers.H, cpu.Registers.D)
}

func (cpu *CPU) LD_H_C() int {
	return cpu.ld_dst_src(cpu.Registers.H, cpu.Registers.C)
}

func (cpu *CPU) LD_D_C() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.C)
}

func (cpu *CPU) LD_D_D() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.D)
}

func (cpu *CPU) LD_D_E() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.E)
}

func (cpu *CPU) LD_D_H() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.H)
}

func (cpu *CPU) LD_D_L() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.L)
}

func (cpu *CPU) LD_D_LOC_HL() int {
	return cpu.ld_dst_loc_src(cpu.Registers.D, cpu.Registers.HL)
}

func (cpu *CPU) LD_D_A() int {
	return cpu.ld_dst_src(cpu.Registers.D, cpu.Registers.A)
}

func (cpu *CPU) LD_E_B() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.B)
}

func (cpu *CPU) LD_E_C() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.C)
}

func (cpu *CPU) LD_E_D() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.D)
}

func (cpu *CPU) LD_E_E() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.E)
}

func (cpu *CPU) LD_E_H() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.H)
}

func (cpu *CPU) LD_E_L() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.L)
}

func (cpu *CPU) LD_E_A() int {
	return cpu.ld_dst_src(cpu.Registers.E, cpu.Registers.A)
}

func (cpu *CPU) LD_E_LOC_HL() int {
	return cpu.ld_dst_loc_src(cpu.Registers.E, cpu.Registers.HL)
}

func (cpu *CPU) LD_B_C() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.C)
}

func (cpu *CPU) LD_B_D() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.D)
}

func (cpu *CPU) LD_B_E() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.E)
}

func (cpu *CPU) LD_B_H() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.H)
}

func (cpu *CPU) LD_B_L() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.L)
}

func (cpu *CPU) LD_B_A() int {
	return cpu.ld_dst_src(cpu.Registers.B, cpu.Registers.A)
}

func (cpu *CPU) LD_C_B() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.B)
}

func (cpu *CPU) LD_C_C() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.C)
}

func (cpu *CPU) LD_C_D() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.D)
}

func (cpu *CPU) LD_C_E() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.E)
}

func (cpu *CPU) LD_C_H() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.H)
}

func (cpu *CPU) LD_C_L() int {
	return cpu.ld_dst_src(cpu.Registers.C, cpu.Registers.L)
}

func (cpu *CPU) LD_B_LOC_HL() int {
	return cpu.ld_dst_loc_src(cpu.Registers.B, cpu.Registers.HL)
}

func (cpu *CPU) LD_C_LOC_HL() int {
	return cpu.ld_dst_loc_src(cpu.Registers.C, cpu.Registers.HL)
}

func (cpu *CPU) LD_A_LOC_BC() int {
	return cpu.ld_dst_loc_src(cpu.Registers.A, cpu.Registers.BC)
}

func (cpu *CPU) ld_dst_d8(register *EightBitRegister) int {
	register.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) LD_H_D8() int {
	return cpu.ld_dst_d8(cpu.Registers.H)
}

func (cpu *CPU) LD_A_LOC_HL_INC() int {
	cpu.Registers.A.Set(cpu.MMU.Read(cpu.Registers.HL.Get()))
	cpu.Registers.HL.Increment()
	return 8
}

func (cpu *CPU) LD_A_LOC_HL_DEC() int {
	cpu.Registers.A.Set(cpu.MMU.Read(cpu.Registers.HL.Get()))
	cpu.Registers.HL.Decrement()
	return 8
}

func (cpu *CPU) LD_LOC_HL_D8() int {
	cpu.MMU.Write(cpu.Registers.HL.Get(), cpu.FetchAndIncrement())
	return 12
}

func (cpu *CPU) LD_A8_A() int {
	cpu.MMU.Write(0xFF00+uint16(cpu.FetchAndIncrement()), cpu.Registers.A.Get())
	return 12
}

func (cpu *CPU) LD_A_A8() int {
	cpu.Registers.A.Set(cpu.MMU.Read(0xFF00 + uint16(cpu.FetchAndIncrement())))
	return 12
}

func (cpu *CPU) LD_A_LOC_C() int {
	cpu.Registers.A.Set(cpu.MMU.Read(0xFF00 + uint16(cpu.Registers.C.Get())))
	return 8
}

func (cpu *CPU) LD_A_LOC_A16() int {
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.Registers.A.Set(cpu.MMU.Read(MergeBytes(high, low)))
	return 16
}

func (cpu *CPU) inc_register(register *EightBitRegister) int {
	register.Increment()
	return 4
}

func (cpu *CPU) dec_register(register *EightBitRegister) int {
	register.Decrement()
	return 4
}

func (cpu *CPU) inc_register_sixteen(register *SixteenBitRegister) int {
	register.Increment()
	return 4
}

func (cpu *CPU) dec_register_sixteen(register *SixteenBitRegister) int {
	register.Decrement()
	return 4
}

func (cpu *CPU) INC_D() int {
	return cpu.inc_register(cpu.Registers.D)
}

func (cpu *CPU) INC_E() int {
	return cpu.inc_register(cpu.Registers.E)
}

func (cpu *CPU) INC_L() int {
	return cpu.inc_register(cpu.Registers.L)
}

func (cpu *CPU) INC_A() int {
	return cpu.inc_register(cpu.Registers.A)
}

func (cpu *CPU) INC_BC() int {
	return cpu.inc_register_sixteen(cpu.Registers.BC)
}

func (cpu *CPU) INC_SP() int {
	cpu.SP += 1
	return 4
}

func (cpu *CPU) DEC_H() int {
	return cpu.dec_register(cpu.Registers.H)
}

func (cpu *CPU) DEC_L() int {
	return cpu.dec_register(cpu.Registers.L)
}

func (cpu *CPU) DEC_SP() int {
	cpu.SP -= 1
	return 4
}

func (cpu *CPU) DEC_BC() int {
	return cpu.dec_register_sixteen(cpu.Registers.BC)
}

func (cpu *CPU) DEC_DE() int {
	return cpu.dec_register_sixteen(cpu.Registers.DE)
}

func (cpu *CPU) DEC_HL() int {
	return cpu.dec_register_sixteen(cpu.Registers.HL)
}

func (cpu *CPU) INC_LOC_HL() int {
	hl := cpu.Registers.HL.Get()
	current := cpu.MMU.Read(hl)
	cpu.MMU.Write(hl, current+1)
	return 12
}

func (cpu *CPU) DEC_LOC_HL() int {
	hl := cpu.Registers.HL.Get()
	current := cpu.MMU.Read(hl)
	cpu.MMU.Write(hl, current-1)
	return 12
}

func (cpu *CPU) add_dst_src_sixteen(dst *SixteenBitRegister, src uint16) int {
	first, second := dst.Get(), src
	sum := first + second
	bytes := SplitInt16ToBytes(sum)
	dst.Set(bytes[0], bytes[1])
	cpu.CheckAndSetOverflowFlagSixteenBit(first,second, false)
	cpu.ClearNegativeFlag()
	cpu.CheckAndSetHCFlagSixteenBit(first,second, false)
	return 8
}

func (cpu *CPU) ADD_HL_BC() int {
	return cpu.add_dst_src_sixteen(cpu.Registers.HL, cpu.Registers.BC.Get())
}

func (cpu *CPU) ADD_HL_DE() int {
	return cpu.add_dst_src_sixteen(cpu.Registers.HL, cpu.Registers.DE.Get())
}

func (cpu *CPU) ADD_HL_HL() int {
	return cpu.add_dst_src_sixteen(cpu.Registers.HL, cpu.Registers.HL.Get())
}

func (cpu *CPU) ADD_HL_SP() int {
	return cpu.add_dst_src_sixteen(cpu.Registers.HL, cpu.SP)
}

func (cpu *CPU) add_dst_src(dst *EightBitRegister, src byte) int {
	first, second := dst.Get(), src
	sum := first + second
	dst.Set(sum)
	cpu.CheckAndSetOverflowFlag(first, second, false)
	cpu.CheckAndSetZeroFlag(sum)
	cpu.ClearNegativeFlag()
	cpu.CheckAndSetHCFlag(first, second, false)
	return 4
}

func (cpu *CPU) sub(register *EightBitRegister) int {
	a := cpu.Registers.A.Get()
	b := register.Get()
	diff := a-b
	cpu.Registers.A.Set(diff)
	cpu.SetNegativeFlag()
	cpu.CheckAndSetZeroFlag(diff)
	cpu.CheckAndSetHCFlag(a, b, true)
	cpu.CheckAndSetOverflowFlag(a, b, true)
	return 4
}

func (cpu *CPU) ADD_A_B() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.B.Get())
}

func (cpu *CPU) ADD_A_C() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.C.Get())
}

func (cpu *CPU) ADD_A_D() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.D.Get())
}

func (cpu *CPU) ADD_A_E() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.E.Get())
}

func (cpu *CPU) ADD_A_H() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.H.Get())
}

func (cpu *CPU) ADD_A_L() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.L.Get())
}

func (cpu *CPU) ADD_A_A() int {
	return cpu.add_dst_src(cpu.Registers.A, cpu.Registers.A.Get())
}

func (cpu *CPU) SUB_B() int {
	return cpu.sub(cpu.Registers.B)
}

func (cpu *CPU) SUB_C() int {
	return cpu.sub(cpu.Registers.C)
}

func (cpu *CPU) SUB_D() int {
	return cpu.sub(cpu.Registers.D)
}

func (cpu *CPU) SUB_E() int {
	return cpu.sub(cpu.Registers.E)
}

func (cpu *CPU) SUB_H() int {
	return cpu.sub(cpu.Registers.H)
}

func (cpu *CPU) SUB_L() int {
	return cpu.sub(cpu.Registers.L)
}

func (cpu *CPU) SUB_A() int {
	return cpu.sub(cpu.Registers.A)
}

func (cpu *CPU) and(register *EightBitRegister) int {
	result := cpu.Registers.A.Get() & register.Get()
	cpu.Registers.A.Set(result)
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.SetHCFlag()
	cpu.ClearOverflowFlag()
	return 4
}

func (cpu *CPU) AND_B() int {
	return cpu.and(cpu.Registers.B)
}

func (cpu *CPU) AND_C() int {
	return cpu.and(cpu.Registers.C)
}

func (cpu *CPU) AND_D() int {
	return cpu.and(cpu.Registers.D)
}

func (cpu *CPU) AND_E() int {
	return cpu.and(cpu.Registers.E)
}

func (cpu *CPU) AND_H() int {
	return cpu.and(cpu.Registers.H)
}

func (cpu *CPU) AND_L() int {
	return cpu.and(cpu.Registers.L)
}

func (cpu *CPU) AND_A() int {
	return cpu.and(cpu.Registers.A)
}

func (cpu *CPU) or(register *EightBitRegister) int {
	result := cpu.Registers.A.Get() | register.Get()
	cpu.Registers.A.Set(result)
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	cpu.ClearOverflowFlag()
	return 4
}

func (cpu *CPU) OR_B() int {
	return cpu.or(cpu.Registers.B)
}

func (cpu *CPU) OR_C() int {
	return cpu.or(cpu.Registers.C)
}

func (cpu *CPU) OR_D() int {
	return cpu.or(cpu.Registers.D)
}

func (cpu *CPU) OR_E() int {
	return cpu.or(cpu.Registers.E)
}

func (cpu *CPU) OR_H() int {
	return cpu.or(cpu.Registers.H)
}

func (cpu *CPU) OR_L() int {
	return cpu.or(cpu.Registers.L)
}

func (cpu *CPU) OR_A() int {
	return cpu.or(cpu.Registers.A)
}


