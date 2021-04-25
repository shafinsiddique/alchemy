package main

func (cpu *CPU) HandleInterrupts() {
	if !cpu.IME {
		return
	}

	for i := 0; i < 5; i++ {
		if cpu.Registers.IF.GetBit(i) == 1 && cpu.Registers.IE.GetBit(i) == 1 {
			cpu.IME = false
			cpu.Registers.IF.SetBit(0, i)
			cpu.handleInterrupt(i)
			return
		}
	}
}

func (cpu *CPU) handleInterrupt(index int) {
	var addr uint16
	switch index {
	case 0:
		addr = 0x40
	case 1:
		addr = 0x48
	case 2:
		addr = 0x50
	case 3:
		addr = 0x58
	case 4:
		addr = 0x60
	}

	cpu.storePcAndJump(addr)
}
