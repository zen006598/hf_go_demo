package enums

type Color byte

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

func (c Color) ToString() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
