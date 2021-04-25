package main

type ReferenceRegister struct {
	Element *byte
}

func NewReferenceRegister(element *byte) *ReferenceRegister {
	return &ReferenceRegister{Element: element}
}

func (register *ReferenceRegister) Get() byte {
	return *register.Element
}

func (register *ReferenceRegister) Set(val byte) {
	*register.Element = val
}

func (register *ReferenceRegister) GetBit(index int) byte {
	val := register.Get()
	return GetBit(val, index)
}

func (register *ReferenceRegister) Increment() {
	*register.Element += 1
}

func (register *ReferenceRegister) Decrement() {
	*register.Element -= 1
}
