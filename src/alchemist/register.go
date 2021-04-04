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
	/*
		index range 0-7. counting from right to left, weird idk.
	*/
	val := register.Get()
	bit := (val & (1 << index)) != 0
	if bit {
		return 1
	}
	return 0
}

func (register *EightBitRegister) SetBit(val byte, index int) {
	var newVal byte
	if val == 1 {
		newVal = val | (1 << index)
	} else {
		newVal = val & ((1 << index) ^ 1) // XOR 1 gets the complement.
	}

	register.Set(newVal)
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
