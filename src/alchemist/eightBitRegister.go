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

func (register *EightBitRegister) Get() byte {
	return register.Value
}

func (register *EightBitRegister) GetBit(index int) byte{
	/* index range 0-7. counting from right to left, weird idk. */
	val := register.Get()
	return GetBit(val, index)
}

func (register *EightBitRegister) SetBit(bit byte, index int) {
	current := register.Get()
	register.Set(SetBit(current, bit, index))
}

func (register *EightBitRegister) Increment(){
	register.Value += 1
}

func (register *EightBitRegister) Decrement() {
	register.Value -= 1
}
