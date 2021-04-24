package color

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var (
	Red   = RGB(255, 0, 0)
	Green = RGB(0, 255, 0)
	Blue  = RGB(0, 0, 255)
	None  = RGBA(0, 0, 0, 0)
)

func RGB(red, green, blue uint8) Color {
	return RGBA(red, green, blue, 255)
}

func RGBA(red, green, blue, alpha uint8) Color {
	return Color{
		R: red,
		G: green,
		B: blue,
		A: alpha,
	}
}

func (c Color) EqualsRGB(other Color) bool {
	return c.R == other.R && c.G == other.G && c.B == other.B
}

func (c Color) EqualsRGBA(other Color) bool {
	return c.EqualsRGB(other) && c.A == other.A
}

func (c Color) Uint32() uint32 {
	return uint32(c.R)<<24 | uint32(c.G)<<16 | uint32(c.B)<<8 | uint32(c.A)
}
