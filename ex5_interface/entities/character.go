package entities

import "math"

type Shape interface {
	Area() float64
}

type Rect struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
