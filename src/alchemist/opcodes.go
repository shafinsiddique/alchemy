package main

func (cpu *CPU) LD_SP_D16() int {
	// 0x31
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.SP = MergeBytes(high, low)
	return 12
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

func (cpu *CPU) jr_s8(bitIndex int, flag byte) int { // not actual instreuction, code reuse.
	zFlag := cpu.Registers.F.GetBit(bitIndex)
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

func (cpu *CPU) RLCA() int {
	cpu.Registers.F.SetBit(cpu.Registers.A.GetBit(7), CARRY_FLAG) // set value of bit 7 to carry flag so we
	// can set bit 0 to it.

	for i := 0; i <= 7; i++ {
		bit := cpu.Registers.A.GetBit(i)
		replace := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.A.SetBit(replace, i)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
	}

	cpu.ClearOverflowFlag()
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()

	return 4
}

func (cpu *CPU) RRA() int {
	for i := 7; i>=0; i-- {
		bit := cpu.Registers.A.GetBit(i)
		replace := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.A.SetBit(replace, i)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
	}
	cpu.ClearOverflowFlag()
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 4
}

func (cpu *CPU) RRCA() int {
	cpu.Registers.F.SetBit(cpu.Registers.A.GetBit(0), CARRY_FLAG) // set value of bit 0 to carry flag so we
	// can set bit 7 to it.

	for i := 7; i >= 0; i-- {
		bit := cpu.Registers.A.GetBit(i)
		replace := cpu.Registers.F.GetBit(CARRY_FLAG)
		cpu.Registers.A.SetBit(replace, i)
		cpu.Registers.F.SetBit(bit, CARRY_FLAG)
	}

	cpu.ClearOverflowFlag()
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()

	return 4

}

func (cpu *CPU) JR_NZ_S8() int {
	return cpu.jr_s8(Z_FLAG, 0)
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
	return cpu.incrementRegister8Bit(cpu.Registers.C)
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

func (cpu *CPU) push(register *SixteenBitRegister) int {
	cpu.PushToStack(register.GetHigh())
	cpu.PushToStack(register.GetLow())
	return 16
}

func (cpu *CPU) PUSH_BC() int {
	// push contents of bc into stack.
	return cpu.push(cpu.Registers.BC)
}

func (cpu *CPU) PUSH_DE() int {
	return cpu.push(cpu.Registers.DE)
}

func (cpu *CPU) PUSH_HL() int {
	return cpu.push(cpu.Registers.HL)
}

func (cpu *CPU) PUSH_AF() int {
	return cpu.push(cpu.Registers.AF)
}

func (cpu *CPU) RLA() int {
	// rotate the contents of reigster a to the left, trhough the carry flag.
	// i.e bit 0 -> bit 1 -> bit 2
	cpu.Registers.F.SetBit(0, Z_FLAG)
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
	cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
	return cpu.RL(cpu.Registers.A)

}

func (cpu *CPU) pop(register *SixteenBitRegister) int {
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	register.Set(high, low)
	return 12
}

func (cpu *CPU) pop_pc() int  { // not instruction
	low := cpu.PopFromStack()
	high := cpu.PopFromStack()
	cpu.PC = MergeBytes(high, low)
	return 16
}

func (cpu *CPU) POP_BC() int {
	// pop the contents from the memory stack into BC.
	return cpu.pop(cpu.Registers.BC)
}

func (cpu *CPU) DEC_B() int {
	return cpu.decrementRegister(cpu.Registers.B)
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
	return cpu.pop_pc()
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
	cpu.cp(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) LD_LOC_A16_A() int {
	// store the contents of register A in the internal ram specified by the 16 bit immeidate operand a16.
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.MMU.Write(MergeBytes(high, low), cpu.Registers.A.Get())
	return LD_SIXTEEN_BIT_CYCLE
}

func (cpu *CPU) DEC_A() int {
	return cpu.decrementRegister(cpu.Registers.A)
}

func (cpu *CPU) JR_Z_S8() int {
	return cpu.jr_s8(Z_FLAG,1)
}

func (cpu *CPU) LD_H_A() int {
	cpu.Registers.H.Set(cpu.Registers.A.Get())
	return 4
}

func (cpu *CPU) INC_B() int {
	return cpu.incrementRegister8Bit(cpu.Registers.B)
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
	return cpu.decrementRegister(cpu.Registers.C)
}

func (cpu *CPU) DEC_E() int {
	return cpu.decrementRegister(cpu.Registers.E)
}

func (cpu *CPU) DEC_D() int {
	return cpu.decrementRegister(cpu.Registers.D)
}

func (cpu *CPU) LD_D_D8() int {
	cpu.Registers.D.Set(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) INC_H() int {
	return cpu.incrementRegister8Bit(cpu.Registers.H)
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

func (cpu *CPU) cp(val byte) int {
	// compare the contents of register A with Val by calculating A-(HL)
	a := cpu.Registers.A.Get()
	cpu.CheckAndSetZeroFlag(a-val)
	cpu.CheckAndSetOverflowFlag(a, val, true)
	cpu.SetNegativeFlag()
	cpu.CheckAndSetHCFlag(a, val, true)
	return 4
}

func (cpu *CPU) CP_LOC_HL() int {
	cpu.cp(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) LD_A_L() int {
	cpu.Registers.A.Set(cpu.Registers.L.Get())
	return 4
}

func (cpu *CPU) LD_A_B() int {
	cpu.Registers.A.Set(cpu.Registers.B.Get())
	return 4
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
	return cpu.incrementRegister8Bit(register)
}

func (cpu *CPU) dec_register(register *EightBitRegister) int {
	return cpu.decrementRegister(register)
}

func (cpu *CPU) inc_register_sixteen(register *SixteenBitRegister) int {
	register.Increment()
	return 8
}

func (cpu *CPU) dec_register_sixteen(register *SixteenBitRegister) int {
	register.Decrement()
	return 8
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

func (cpu *CPU) inc_or_dec_loc(register *SixteenBitRegister, dec bool) int {
	loc := register.Get()
	initial := cpu.MMU.Read(loc)
	var val byte
	if dec {
		val = initial-1
		cpu.SetNegativeFlag()
	} else {
		val = initial + 1
		cpu.ClearNegativeFlag()
	}
	cpu.MMU.Write(loc, val)
	cpu.CheckAndSetZeroFlag(val)
	cpu.CheckAndSetHCFlag(initial, val, dec)
	return 12
}

func (cpu *CPU) INC_LOC_HL() int {
	return cpu.inc_or_dec_loc(cpu.Registers.HL, false)
}

func (cpu *CPU) DEC_LOC_HL() int {
	return cpu.inc_or_dec_loc(cpu.Registers.HL, true)
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

func (cpu *CPU) sub(val byte) int {
	a := cpu.Registers.A.Get()
	b := val
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
	return cpu.sub(cpu.Registers.B.Get())
}

func (cpu *CPU) SUB_C() int {
	return cpu.sub(cpu.Registers.C.Get())
}

func (cpu *CPU) SUB_D() int {
	return cpu.sub(cpu.Registers.D.Get())
}

func (cpu *CPU) SUB_E() int {
	return cpu.sub(cpu.Registers.E.Get())
}

func (cpu *CPU) SUB_H() int {
	return cpu.sub(cpu.Registers.H.Get())
}

func (cpu *CPU) SUB_L() int {
	return cpu.sub(cpu.Registers.L.Get())
}

func (cpu *CPU) SUB_A() int {
	return cpu.sub(cpu.Registers.A.Get())
}

func (cpu *CPU) and(val byte) int {
	result := cpu.Registers.A.Get() & val
	cpu.Registers.A.Set(result)
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.SetHCFlag()
	cpu.ClearOverflowFlag()
	return 4
}

func (cpu *CPU) AND_B() int {
	return cpu.and(cpu.Registers.B.Get())
}

func (cpu *CPU) AND_C() int {
	return cpu.and(cpu.Registers.C.Get())
}

func (cpu *CPU) AND_D() int {
	return cpu.and(cpu.Registers.D.Get())
}

func (cpu *CPU) AND_E() int {
	return cpu.and(cpu.Registers.E.Get())
}

func (cpu *CPU) AND_H() int {
	return cpu.and(cpu.Registers.H.Get())
}

func (cpu *CPU) AND_L() int {
	return cpu.and(cpu.Registers.L.Get())
}

func (cpu *CPU) AND_A() int {
	return cpu.and(cpu.Registers.A.Get())
}

func (cpu *CPU) or(val byte) int {
	result := cpu.Registers.A.Get() | val
	cpu.Registers.A.Set(result)
	cpu.CheckAndSetZeroFlag(result)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	cpu.ClearOverflowFlag()
	return 4
}

func (cpu *CPU) OR_B() int {
	return cpu.or(cpu.Registers.B.Get())
}

func (cpu *CPU) OR_C() int {
	return cpu.or(cpu.Registers.C.Get())
}

func (cpu *CPU) OR_D() int {
	return cpu.or(cpu.Registers.D.Get())
}

func (cpu *CPU) OR_E() int {
	return cpu.or(cpu.Registers.E.Get())
}

func (cpu *CPU) OR_H() int {
	return cpu.or(cpu.Registers.H.Get())
}

func (cpu *CPU) OR_L() int {
	return cpu.or(cpu.Registers.L.Get())
}

func (cpu *CPU) OR_A() int {
	return cpu.or(cpu.Registers.A.Get())
}

func (cpu *CPU) ADD_A_LOC_HL() int {
	cpu.add_dst_src(cpu.Registers.A, cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) SUB_LOC_HL() int {
	cpu.sub(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) AND_LOC_HL() int {
	cpu.and(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) OR_LOC_HL() int {
	cpu.or(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) ADD_A_D8() int {
	cpu.add_dst_src(cpu.Registers.A, cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) SUB_D8() int {
	cpu.sub(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) AND_D8() int {
	cpu.and(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) OR_D8() int {
	cpu.or(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) xor(val byte) int {
	cpu.Registers.A.Set(val ^ val)
	cpu.CheckAndSetZeroFlag(cpu.Registers.A.Get())
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	cpu.ClearOverflowFlag()
	return 4
}

func (cpu *CPU) XOR_A() int {
	return cpu.xor(cpu.Registers.A.Get())
}

func (cpu *CPU) XOR_B() int {
	return cpu.xor(cpu.Registers.B.Get())
}

func (cpu *CPU) XOR_C() int {
	return cpu.xor(cpu.Registers.C.Get())
}

func (cpu *CPU) XOR_D() int {
	return cpu.xor(cpu.Registers.D.Get())
}

func (cpu *CPU) XOR_E() int {
	return cpu.xor(cpu.Registers.E.Get())
}

func (cpu *CPU) XOR_H() int {
	return cpu.xor(cpu.Registers.H.Get())
}

func (cpu *CPU) XOR_L() int {
	return cpu.xor(cpu.Registers.L.Get())
}

func (cpu *CPU) XOR_LOC_HL() int {
	cpu.xor(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) XOR_D8() int {
	cpu.xor(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) CP_B() int {
	return cpu.cp(cpu.Registers.B.Get())
}

func (cpu *CPU) CP_C() int {
	return cpu.cp(cpu.Registers.C.Get())
}

func (cpu *CPU) CP_D() int {
	return cpu.cp(cpu.Registers.D.Get())
}

func (cpu *CPU) CP_E() int {
	return cpu.cp(cpu.Registers.E.Get())
}

func (cpu *CPU) CP_H() int {
	return cpu.cp(cpu.Registers.H.Get())
}

func (cpu *CPU) CP_L() int {
	return cpu.cp(cpu.Registers.L.Get())
}

func (cpu *CPU) CP_A() int {
	return cpu.cp(cpu.Registers.A.Get())
}

func (cpu *CPU) POP_DE() int {
	return cpu.pop(cpu.Registers.DE)
}

func (cpu *CPU) POP_HL() int {
	return cpu.pop(cpu.Registers.HL)
}

func (cpu *CPU) POP_AF() int {
	return cpu.pop(cpu.Registers.AF)
}

func (cpu *CPU) LD_HL_SP_S8() int {
	s8, isNegative := GetTwosComplement(cpu.FetchAndIncrement())
	val := cpu.SP
	if isNegative {
		val -= uint16(s8)
	} else {
		val += uint16(s8)
	}

	cpu.Registers.HL.SetV2(val)
	return 12
}

func (cpu *CPU) LD_SP_HL() int {
	cpu.SP = cpu.Registers.HL.Get()
	return 8
}

func (cpu *CPU) ADD_SP_S8() int {
	current := cpu.SP
	val, isNegative := GetTwosComplement(cpu.FetchAndIncrement())
	val16 := uint16(val)
	if isNegative {
		current -= val16
	} else {
		current += val16
	}

	cpu.SP = current
	cpu.ClearZeroFlag()
	cpu.ClearNegativeFlag()
	cpu.CheckAndSetHCFlagSixteenBit(current, val16, isNegative)
	cpu.CheckAndSetOverflowFlagSixteenBit(current, val16, isNegative)
	return 16
}

func (cpu *CPU) adc(val byte) int {
	cy := cpu.Registers.F.GetBit(CARRY_FLAG)
	initial := cpu.Registers.A.Get()
	sum := val + cy + initial
	cpu.Registers.A.Set(sum)
	cpu.CheckAndSetZeroFlag(sum)
	cpu.ClearNegativeFlag()
	if !cpu.CheckAndSetOverflowFlag(val, cy, false)  {
		cpu.CheckAndSetOverflowFlag(val+cy, initial, false)
	}
	cpu.CheckAndSetHCFlag(initial, val+cy, false)
	return 8
}

func (cpu *CPU) ADC_A_B() int {
	return cpu.adc(cpu.Registers.B.Get())
}
func (cpu *CPU) ADC_A_C() int {
	return cpu.adc(cpu.Registers.C.Get())
}
func (cpu *CPU) ADC_A_D() int {
	return cpu.adc(cpu.Registers.D.Get())
}
func (cpu *CPU) ADC_A_E() int {
	return cpu.adc(cpu.Registers.E.Get())
}
func (cpu *CPU) ADC_A_H() int {
	return cpu.adc(cpu.Registers.H.Get())
}
func (cpu *CPU) ADC_A_L() int {
	return cpu.adc(cpu.Registers.L.Get())
}
func (cpu *CPU) ADC_A_A() int {
	return cpu.adc(cpu.Registers.A.Get())
}

func (cpu *CPU) ADC_A_LOC_HL() int {
	cpu.adc(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) sbc(val byte) int {
	initial := cpu.Registers.A.Get()
	cy := cpu.Registers.F.GetBit(CARRY_FLAG)
	sub := val + cy
	result := initial - sub
	cpu.Registers.A.Set(result)
	cpu.CheckAndSetZeroFlag(result)
	cpu.SetNegativeFlag()
	if !cpu.CheckAndSetOverflowFlag(val, cy, false) {
		cpu.CheckAndSetOverflowFlag(initial, sub, true)
	}

	cpu.CheckAndSetHCFlag(initial, sub, true)
	return 4
}

func (cpu *CPU) SBC_A_B() int {
	return cpu.sbc(cpu.Registers.B.Get())
}

func (cpu *CPU) SBC_A_C() int {
	return cpu.sbc(cpu.Registers.C.Get())
}

func (cpu *CPU) SBC_A_D() int {
	return cpu.sbc(cpu.Registers.D.Get())
}

func (cpu *CPU) SBC_A_E() int {
	return cpu.sbc(cpu.Registers.E.Get())
}

func (cpu *CPU) SBC_A_H() int {
	return cpu.sbc(cpu.Registers.H.Get())
}

func (cpu *CPU) SBC_A_L() int {
	return cpu.sbc(cpu.Registers.L.Get())
}

func (cpu *CPU) SBC_A_A() int {
	return cpu.sbc(cpu.Registers.A.Get())
}

func (cpu *CPU) SBC_A_LOC_HL() int {
	cpu.sbc(cpu.MMU.Read(cpu.Registers.HL.Get()))
	return 8
}

func (cpu *CPU) SBC_A_D8() int {
	cpu.sbc(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) ADC_A_D8() int {
	cpu.adc(cpu.FetchAndIncrement())
	return 8
}

func (cpu *CPU) JR_NC_S8() int {
	return cpu.jr_s8(CARRY_FLAG, 0)
}

func (cpu *CPU) JR_C_S8() int {
	return cpu.jr_s8(CARRY_FLAG, 1)
}

func (cpu *CPU) ret_flag(flagIndex int, bit byte) int {
	flag := cpu.Registers.F.GetBit(flagIndex)
	if flag == bit {
		cpu.pop_pc()
		return 20
	}
	return 8
}

func (cpu *CPU) RET_NZ() int {
	return cpu.ret_flag(Z_FLAG, 0)
}

func (cpu *CPU) RET_NC() int {
	return cpu.ret_flag(CARRY_FLAG, 0)
}

func (cpu *CPU) RET_Z() int {
	return cpu.ret_flag(Z_FLAG, 1)
}

func (cpu *CPU) RET_C() int {
	return cpu.ret_flag(CARRY_FLAG, 1)
}

func (cpu *CPU) RETI() int {
	cpu.IME = true
	return cpu.RET()
}

func (cpu *CPU) PUSH_16(val uint16)  {
	bytes := SplitInt16ToBytes(val)
	cpu.PushToStack(bytes[0])
	cpu.PushToStack(bytes[1])
}

func (cpu *CPU) LD_LOC_A16_SP() int {
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	bytes := SplitInt16ToBytes(cpu.SP)
	loc := MergeBytes(high, low)
	cpu.MMU.Write(loc, bytes[1])
	cpu.MMU.Write(loc+1, bytes[0])
	return 20
}

func (cpu *CPU) STOP() int {
	return 4
}

func (cpu *CPU) DAA() int {
	return 4
}

func (cpu *CPU) CPL() int {
	cpu.Registers.A.Set(GetOnesComplement(cpu.Registers.A.Get()))
	cpu.SetNegativeFlag()
	cpu.SetHCFlag()
	return 4
}

func (cpu *CPU) SCF() int {
	cpu.SetOverflowFlag()
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 4
}

func (cpu *CPU) CCF() int {
	cy := cpu.Registers.F.GetBit(CARRY_FLAG)
	cy = 1-cy // probs could have been done with bit manipualtion.
	cpu.Registers.F.SetBit(cy, CARRY_FLAG)
	cpu.ClearNegativeFlag()
	cpu.ClearHCFlag()
	return 4
}

func (cpu *CPU) interruptExists() bool {
	for i := 0; i < 5; i++ {
		if cpu.Registers.IF.GetBit(i) == 1 && cpu.Registers.IE.GetBit(i) == 1 {
			return true
		}
	}
	return false
}

func (cpu *CPU) HALT() int {
	var decrementPc bool
	if cpu.Halted {
		if cpu.interruptExists() {
			cpu.Halted = false
		} else {
			decrementPc = true
		}
	} else {
		decrementPc = true
		cpu.Halted = true
	}

	if decrementPc {
		pc := &cpu.PC
		*pc -= 1
	}

	return 4
}

func (cpu *CPU) rst(index uint16) int {
	cpu.PUSH_16(cpu.PC)
	cpu.PC = 0x0000 + (8*index)
	return 16
}

func (cpu *CPU) RST_0() int {
	return cpu.rst(0)
}

func (cpu *CPU) RST_1() int {
	return cpu.rst(1)
}

func (cpu *CPU) RST_2() int {
	return cpu.rst(2)
}

func (cpu *CPU) RST_3() int {
	return cpu.rst(3)
}

func (cpu *CPU) RST_4() int {
	return cpu.rst(4)
}

func (cpu *CPU) RST_5() int {
	return cpu.rst(5)
}

func (cpu *CPU) RST_6() int {
	return cpu.rst(6)
}

func (cpu *CPU) RST_7() int {
	return cpu.rst(7)
}

func (cpu *CPU) jp_a16_conditonal(flagIndex int, flagValue byte) int  {
	flag := cpu.Registers.F.GetBit(flagIndex)
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cycles := 12
	if flag == flagValue  {
		cpu.PC = MergeBytes(high, low)
		cycles = 16
	}

	return cycles
}

func (cpu *CPU) JP_A16() int {
	low := cpu.FetchAndIncrement()
	high := cpu.FetchAndIncrement()
	cpu.PC = MergeBytes(high, low)
	return 16
}

func (cpu *CPU) DI() int {
	cpu.IME = false
	return 4
}

func (cpu *CPU) EI() int {
	cpu.IME = true
	return 4
}

func (cpu *CPU) JP_NZ_A16() int {
	return cpu.jp_a16_conditonal(Z_FLAG, 0)
}

func (cpu *CPU) JP_Z_A16() int {
	return cpu.jp_a16_conditonal(Z_FLAG, 1)
}

func (cpu *CPU) JP_NC_A16() int {
	return cpu.jp_a16_conditonal(CARRY_FLAG, 0)
}

func (cpu *CPU) JP_C_A16() int {
	return cpu.jp_a16_conditonal(CARRY_FLAG, 1)
}

func (cpu *CPU) call_a16(flagIndex int, flagValue byte) int {
	flag := cpu.Registers.F.GetBit(flagIndex)
	low, high := cpu.FetchAndIncrement(), cpu.FetchAndIncrement()
	cycles := 12
	if flag == flagValue {
		cpu.PUSH_16(cpu.PC)
		cpu.PC = MergeBytes(high, low)
		cycles = 24
	}
	return cycles
}

func (cpu *CPU) CALL_NZ_A16() int {
	return cpu.call_a16(Z_FLAG, 0)
}

func (cpu *CPU) CALL_Z_A16() int {
	return cpu.call_a16(Z_FLAG, 1)
}

func (cpu *CPU) CALL_NC_A16() int {
	return cpu.call_a16(CARRY_FLAG, 0)
}

func (cpu *CPU) CALL_C_A16() int {
	return cpu.call_a16(CARRY_FLAG, 1)
}

func (cpu *CPU) JP_HL() int {
	cpu.PC = cpu.Registers.HL.Get()
	return 4
}

func (cpu *CPU) NOP() int {
	return 4
}