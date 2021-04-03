package main

func MergeBytes(high byte, low byte) uint16{
	high_16 := uint16(high)
	low_16 := uint16(low)

	return high_16 << 8 | low_16
}