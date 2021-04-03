package main

import "fmt"

func main() {
	registers := InitializeRegisters()
	registers.A.Set(0xA2)
	registers.F.Set(0xF0)
	fmt.Println(registers.AF.Get())
	registers.A.Set(0xC2)
	fmt.Println(registers.AF.Get())

}
