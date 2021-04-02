package main

type EightBitRegister struct {
	Value byte
}


func (register *EightBitRegister) Set(value byte) {
	register.Value = value
}

type SixteenBitRegister struct {
	LowByte *EightBitRegister
	HighByte *EightBitRegister
}
