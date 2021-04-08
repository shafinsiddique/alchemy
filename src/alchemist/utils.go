package main

import (
	"encoding/binary"
)

func MergeBytes(high byte, low byte) uint16{
	high_16 := uint16(high)
	low_16 := uint16(low)

	return high_16 << 8 | low_16
}

func SplitInt16ToBytes(num uint16)[]byte {
	output := make([]byte, 2)
	binary.BigEndian.PutUint16(output, num)
	return output
}

func GetBit(val byte, index int) byte {
	bit := (val & (1 << index)) != 0

	if bit {
		return 1
	}

	return 0
}

func SetBit(value byte, bit byte, index int) byte {
	var newVal byte
	if bit == 1 {
		newVal = value | (1 << index)
	} else {
		newVal = value & (^(1 << index)) // XOR 1 gets the complement.
	}
	return newVal
}

func GetTwosComplement(value byte) (byte, bool){ // using bool to indicate whether a number is negative.
	// probably a better way to do this but i don't want the program counter to be a SIGNED int.
	sign := GetBit(value, 7)
	val := value
	if sign == 1 {
		val = ^val
		val += 1
		return val, true // if true, number is negative.
	}
	return val, false
}

