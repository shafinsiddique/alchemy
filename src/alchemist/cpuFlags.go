package main

func (cpu *CPU) SetZeroFlag() {
	cpu.Registers.F.SetBit(1, Z_FLAG)
}

func (cpu *CPU) ClearZeroFlag() {
	cpu.Registers.F.SetBit(0, Z_FLAG)
}

func (cpu *CPU) SetNegativeFlag() {
	cpu.Registers.F.SetBit(1, NEGATIVE_FLAG)
}

func (cpu *CPU) ClearNegativeFlag() {
	cpu.Registers.F.SetBit(0, NEGATIVE_FLAG)
}

func (cpu *CPU) CheckAndSetZeroFlag(value byte) bool {
	zero := false
	if value == 0 {
		cpu.Registers.F.SetBit(1, Z_FLAG)
		zero = true
	} else {
		cpu.Registers.F.SetBit(0, Z_FLAG)
	}
	return zero
}

func (cpu *CPU) CheckAndSetZeroFlagSixteenBit(value uint16) bool {
	zero := false
	if value == 0 {
		cpu.Registers.F.SetBit(1, Z_FLAG)
		zero = true
	} else {
		cpu.Registers.F.SetBit(0, Z_FLAG)
	}
	return zero
}

func checkHCFlag(a byte, b byte, negative bool) bool {
	var sum byte
	carry := false
	if negative {
		sum = (a & 0xf) - (b & 0xf)
	} else {
		sum = (a & 0xf) + (b & 0xf)
	}
	if (sum & 0x10) == 0x10 {
		carry = true
	}

	return carry
}

func (cpu *CPU) CheckAndSetHCFlag(a byte, b byte, negative bool) bool {
	carry := checkHCFlag(a, b, negative)
	if carry {
		cpu.Registers.F.SetBit(1, HALF_CARRY_FLAG)
	} else {
		cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
	}
	return carry
}

func (cpu *CPU) CheckAndSetHCFlagSixteenBit(a uint16, b uint16, negative bool) bool { // Code repetition, fix later.
	originalBytes := SplitInt16ToBytes(a)
	currentBytes := SplitInt16ToBytes(b)

	carry := checkHCFlag(originalBytes[0], currentBytes[0], negative) // check hc on high byte.

	if carry {
		cpu.Registers.F.SetBit(1, HALF_CARRY_FLAG)
	} else {
		cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
	}

	return carry
}

func (cpu *CPU) CheckAndSetOverflowFlag(a byte, b byte, negative bool) bool {
	var overflow bool
	if negative {
		if a < b+0 {
			overflow = true
		}
	} else {
		if a > 255-b {
			overflow = true
		}
	}

	if overflow {
		cpu.Registers.F.SetBit(1, CARRY_FLAG)
	} else {
		cpu.Registers.F.SetBit(0, CARRY_FLAG)
	}

	return overflow
}

func (cpu *CPU) CheckAndSetOverflowFlagSixteenBit(a uint16, b uint16, negative bool) bool {
	var overflow bool
	if negative {
		if a < b+0 {
			overflow = true
		}
	} else {
		if a > 65535-b {
			overflow = true
		}
	}

	if overflow {
		cpu.Registers.F.SetBit(1, CARRY_FLAG)
	} else {
		cpu.Registers.F.SetBit(0, CARRY_FLAG)
	}

	return overflow
}

func (cpu *CPU) SetHCFlag() {
	cpu.Registers.F.SetBit(1, HALF_CARRY_FLAG)
}

func (cpu *CPU) ClearHCFlag() {
	cpu.Registers.F.SetBit(0, HALF_CARRY_FLAG)
}

func (cpu *CPU) ClearOverflowFlag() {
	cpu.Registers.F.SetBit(0, CARRY_FLAG)
}

func (cpu *CPU) SetOverflowFlag() {
	cpu.Registers.F.SetBit(1, CARRY_FLAG)
}
