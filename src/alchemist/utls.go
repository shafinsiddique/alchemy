package main

func MergeBytes(high byte, low byte) uint16{
	high_16 := uint16(high)
	low_16 := uint16(low)

	return high_16 << 8 | low_16
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

func GetTwosComplement(value byte) int{
	sign := GetBit(value, 7)
	val := int8(value)
	if sign == 1 {
		val = ^val
		val += 1
		val = 0-val
	}
	return int(val)
}