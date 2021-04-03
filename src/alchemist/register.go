package main

import (
	"encoding/binary"
	"fmt"
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
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, val)
	fmt.Println("here")
	fmt.Println(bytes[0])
	fmt.Println(bytes[1])
	fmt.Println(MergeBytes(bytes[0],bytes[1]))
	fmt.Println(val)
}
