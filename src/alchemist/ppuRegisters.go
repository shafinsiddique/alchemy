package main

type PPURegisters struct {
	LCDC *PPURegister
	LY   *PPURegister
}

func InitializePPURegisters(memory []byte) *PPURegisters {
	ly := &PPURegister{&memory[LY_INDEX]}
	lcdc := &PPURegister{&memory[LCDC_INDEX]}
	return &PPURegisters{LY: ly, LCDC: lcdc}
}
