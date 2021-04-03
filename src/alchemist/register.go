package main

type EightBitRegister struct {
	Value byte
}

func NewEightBitRegister() *EightBitRegister {
	return &EightBitRegister{}
}

func (register *EightBitRegister) Set(value byte) {
	register.Value = value
}

type SixteenBitRegister struct {
	Low *EightBitRegister
	High *EightBitRegister
}

func NewSixteenBitRegister(high *EightBitRegister, low *EightBitRegister,)*SixteenBitRegister {
	return &SixteenBitRegister{Low: low, High: high}
}

func (register *SixteenBitRegister) Get()uint16 {
	high := uint16(register.High.Value)
	low := uint16(register.Low.Value)

	return high << 8 | low
}
