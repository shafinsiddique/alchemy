package main

type PPURegisters struct {
	LCDC *PPURegister
	LY   *PPURegister
	SCY  *PPURegister
	SCX  *PPURegister
}

func InitializePPURegisters(memory []byte) *PPURegisters {
	ly := &PPURegister{&memory[LY_INDEX]}
	lcdc := &PPURegister{&memory[LCDC_INDEX]}
	scy := &PPURegister{&memory[SCY_INDEX]}
	scx := &PPURegister{&memory[SCX_INDEX]}
	return &PPURegisters{LY: ly, LCDC: lcdc, SCY: scy, SCX: scx}
}
