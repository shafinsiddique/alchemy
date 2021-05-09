package main

func (cpu *CPU) UpdateTimers(cycles byte) {
	cpu.resetIfTimerChanged()

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

func (cpu *CPU) requestTimerInterrupt() {
	cpu.Registers.IF.SetBit(1, TIMER)
}

func (cpu *CPU) timerIsEnabled() bool {
	return cpu.Registers.TAC.GetBit(2) == 1
}

func (cpu *CPU) getTimer() (timer int) {
	val := cpu.Registers.TAC.Get() & 0x3
	switch val {
	case 0:
		timer = 1024
	case 1:
		timer = 16
	case 2:
		timer = 64
	case 3:
		timer = 256
	}

	return timer
}

func (cpu *CPU) resetIfTimerChanged() {
	current := cpu.getTimer()
	existing := cpu.currentTimer

	if current != existing {
		cpu.setTimer()
	}
}

func (cpu *CPU) setTimer() {
	timer := cpu.getTimer()
	cpu.Timer = timer
	cpu.currentTimer = timer
}