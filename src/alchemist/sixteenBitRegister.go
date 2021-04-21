package main

type SixteenBitRegister struct {
	Low  *EightBitRegister
	High *EightBitRegister
}

func NewSixteenBitRegister(high *EightBitRegister, low *EightBitRegister) *SixteenBitRegister {
	return &SixteenBitRegister{Low: low, High: high}
}

func (register *SixteenBitRegister) Get() uint16 {
	high := uint16(register.High.Value)
	low := uint16(register.Low.Value)

	return high<<8 | low
}

func (register *SixteenBitRegister) Set(high byte, low byte) {
	register.Low.Set(high)
	register.High.Set(low)
}

func (register *SixteenBitRegister) Decrement() int {
	val := register.Get()
	decremented := val - 1
	bytes := SplitInt16ToBytes(decremented)
	return register.High.Set(bytes[0]) + register.Low.Set(bytes[1])
}

func (register *SixteenBitRegister) Increment() int {
	val := register.Get()
	incremented := val + 1
	bytes := SplitInt16ToBytes(incremented)
	return register.High.Set(bytes[0]) + register.Low.Set(bytes[1])
}

func (register *SixteenBitRegister) GetHigh() byte {
	return register.High.Get()
}

func (register *SixteenBitRegister) GetLow() byte {
	return register.Low.Get()
}
