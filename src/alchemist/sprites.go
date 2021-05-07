package main

type Sprite struct {
	Y byte
	X byte
	TileID byte
	ObjToBG byte
	YFlip bool
	XFlip bool
	Palette byte
}


func NewSprite(byte0 byte, byte1 byte, byte2 byte, byte3 byte) *Sprite {
	y := byte0
	x := byte1
	tile := byte2
	objToBg := GetBit(byte3, 7)
	yFlip :=  GetBit(byte3, 6) == 1
	xFlip := GetBit(byte3, 5) == 1
	palette := GetBit(byte3, 4)
	return &Sprite{X: x, Y: y, TileID: tile, ObjToBG: objToBg, YFlip: yFlip, XFlip: xFlip, Palette: palette}
}
