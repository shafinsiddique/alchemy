package main

type Registers struct {
	A  *EightBitRegister
	B  *EightBitRegister
	C  *EightBitRegister
	D  *EightBitRegister
	E  *EightBitRegister
	F  *EightBitRegister
	H  *EightBitRegister
	L  *EightBitRegister
	AF *SixteenBitRegister
	BC *SixteenBitRegister
	DE *SixteenBitRegister
	HL *SixteenBitRegister
	IE *ReferenceRegister
	IF *ReferenceRegister
	TIMA *ReferenceRegister
	TMA *ReferenceRegister
	TAC *ReferenceRegister
	DIV *ReferenceRegister
}

func InitializeRegisters(memory []byte) *Registers {
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
	ie := NewReferenceRegister(&memory[IE_INDEX])
	_if := NewReferenceRegister(&memory[IF_INDEX])
	tima := NewReferenceRegister(&memory[TIMA_INDEX])
	tma := NewReferenceRegister(&memory[TMA_INDEX])
	tmc := NewReferenceRegister(&memory[TAC_INDEX])
	div := NewReferenceRegister(&memory[DIV_INDEX])
	return &Registers{A: a, B: b, C: c, D: d, E: e, F: f, H: h, L: l, AF: af, BC: bc, DE: de,
		HL: hl, IE: ie, IF: _if, TIMA: tima, TMA: tma, TAC: tmc, DIV: div}
}
