package main

type IRegister interface {
	Increment()
	Decrement()
	Get() byte
	Set(val byte)
}
