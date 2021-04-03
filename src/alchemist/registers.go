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

func InitializeRegisters() *Registers {
	a := NewEightBitRegister()
	b := NewEightBitRegister()
	c := NewEightBitRegister()
	d := NewEightBitRegister()
	e := NewEightBitRegister()
	f := NewEightBitRegister()
	h := NewEightBitRegister()
	l := NewEightBitRegister()
	af := NewSixteenBitRegister(a, f)
	bc := NewSixteenBitRegister(b, c)
	de := NewSixteenBitRegister(d, e)
	hl := NewSixteenBitRegister(h, l)

	return &Registers{A: a, B: b, C: c, D: d, E: e, F: f, H: h, L:l , AF:af, BC: bc, DE:de,
		HL:hl}
}
