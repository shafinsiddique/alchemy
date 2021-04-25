package main

type PPURegisters struct {
	LCDC *ReferenceRegister
	LY   *ReferenceRegister
	SCY  *ReferenceRegister
	SCX  *ReferenceRegister
	BGP  *ReferenceRegister
	IF *ReferenceRegister
}

func InitializePPURegisters(memory []byte) *PPURegisters {
	ly := &ReferenceRegister{&memory[LY_INDEX]}
	lcdc := &ReferenceRegister{&memory[LCDC_INDEX]}
	scy := &ReferenceRegister{&memory[SCY_INDEX]}
	scx := &ReferenceRegister{&memory[SCX_INDEX]}
	bgp := &ReferenceRegister{&memory[BGP_INDEX]}
	_if := &ReferenceRegister{&memory[IF_INDEX]}
	return &PPURegisters{LY: ly, LCDC: lcdc, SCY: scy, SCX: scx, BGP: bgp, IF:_if}
}
