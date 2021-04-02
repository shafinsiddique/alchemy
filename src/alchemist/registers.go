package main

type Registers struct {
	A *EightBitRegister
	B *EightBitRegister
	C *EightBitRegister
	D *EightBitRegister
	E *EightBitRegister
	F *EightBitRegister
	H *EightBitRegister
	L *EightBitRegister
	AF *SixteenBitRegister
	BC *SixteenBitRegister
	DE *SixteenBitRegister
	HL *SixteenBitRegister
}
