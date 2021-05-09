package main

func (cpu *CPU) UpdateTimers(cycles byte) {
	if !cpu.timerIsEnabled() {
		return
	}

	cpu.Timer -= int(cycles)

	if cpu.Timer <= 0 {
		if cpu.Counter == 255 {
			cpu.Counter = 0
		} else {
			cpu.Counter += 1
		}

		cpu.setTimer()

	}
}

func (cpu *CPU) timerIsEnabled() bool {
	return cpu.Registers.TAC.GetBit(2) == 1
}


func (cpu *CPU) getTimer() (frequency int) {
	val := cpu.Registers.TAC.Get() & 0x3
	switch val {
	case 0:
		frequency = 1024
	case 1:
		frequency = 16
	case 2:
		frequency = 64
	case 3:
		frequency = 256
	}

	return frequency
}

func (cpu *CPU) setTimer() {
	timer := cpu.getTimer()
	cpu.Timer = timer
	cpu.currentTimer = timer
}