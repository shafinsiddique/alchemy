package main

type PPURegisters struct {
	LCDC *ReferenceRegister
	LY   *ReferenceRegister
	SCY  *ReferenceRegister
	SCX  *ReferenceRegister
	BGP  *ReferenceRegister
	IF *ReferenceRegister
	LCD_STAT *ReferenceRegister
	OBP0 *ReferenceRegister
	OBP1 *ReferenceRegister
}

func InitializePPURegisters(memory []byte) *PPURegisters {
	ly := &ReferenceRegister{&memory[LY_INDEX]}
	lcdc := &ReferenceRegister{&memory[LCDC_INDEX]}
	scy := &ReferenceRegister{&memory[SCY_INDEX]}
	scx := &ReferenceRegister{&memory[SCX_INDEX]}
	bgp := &ReferenceRegister{&memory[BGP_INDEX]}
	_if := &ReferenceRegister{&memory[IF_INDEX]}
	lcdStat := &ReferenceRegister{&memory[LCD_STATUS]}
	obp0 := &ReferenceRegister{&memory[OBP0_INDEX]}
	obp1 := &ReferenceRegister{&memory[OBP1_INDEX]}

	return &PPURegisters{LY: ly, LCDC: lcdc, SCY: scy, SCX: scx, BGP: bgp, IF:_if,
		LCD_STAT: lcdStat, OBP0: obp0, OBP1: obp1}
}
