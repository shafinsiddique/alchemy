package main

type PPURegister struct {
	Element *byte
}

func (register *PPURegister) Get() byte{
	return *register.Element
}

func (register *PPURegister) Set(val byte){
	*register.Element = val
}

func (register *PPURegister) GetBit(index int) byte {
	val := register.Get()
	return GetBit(val, index)
}

func (register *PPURegister) Increment(){
	*register.Element += 1
}

func (register *PPURegister) Decrement() {
	*register.Element -= 1
}


