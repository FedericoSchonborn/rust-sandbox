package color

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

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
