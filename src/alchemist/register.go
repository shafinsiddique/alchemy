package main

import (
	"encoding/binary"
)

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

func (register *SixteenBitRegister) Decrement(){
	val := register.Get()
	decremented := val-1
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, decremented)
	register.High.Set(bytes[0])
	register.Low.Set(bytes[1])
}

func (register *SixteenBitRegister) GetHigh() byte{
	return register.High.Get()
}

func (register *SixteenBitRegister) GetLow() byte {
	return register.Low.Get()
}
