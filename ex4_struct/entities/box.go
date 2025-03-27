package entity

import "github.com/zen006598/goDemo/ex4/enums"

type Box struct {
	Width  int
	Height int
	Depth  int
	Color  enums.Color
}

type Boxes []Box

func (b Box) Volume() int {
	return b.Width * b.Height * b.Depth
}

func (b *Box) SetColor(c enums.Color) {
	b.Color = c
}
